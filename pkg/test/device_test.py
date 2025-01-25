import pytest
import requests

device_id = None
data = {
	"Name": "MeasurementIoT",
	"Type": "Sensor",
	"Location": "Field A",
	"Status": "Active",
	"Description": "This measure humidity and temperature",
	"Owner": None
}

dataUser = {
    "Name":"Alpha Zero One",
    "Email":"alpha01@mail.com",
    "Password":"password"
}

def test_create_device():
    """Test the POST /devices endpoint."""
    global device_id

    data["Owner"] = requests.post("http://127.0.0.1:8888/users", json=dataUser).json()["data"]["User_Id"]

    response = requests.post("http://127.0.0.1:8888/devices", json=data)
    assert response.status_code == 201 

    response = response.json()["data"]
    assert response["Name"] == data["Name"]
    assert response["Type"] == data["Type"]
    assert response["Location"] == data["Location"]
    assert response["Status"] == data["Status"]
    assert response["Description"] == data["Description"]
    assert "Device_Id" in response
    assert "Token" in response
    assert "Created_At" in response

    device_id = response["Device_Id"]

def test_create_device_invalid_data():
    """Test the POST /users endpoint with invalid data."""

    invalidData = {
        "Name": "This is Invalid Example" 
    }

    response = requests.post("http://127.0.0.1:8888/devices", json=invalidData)
    assert response.status_code == 400

    response = response.json()["errors"]
    assert "timestamp" in response
    assert "status" in response
    assert "error" in response
    assert "message" in response
    assert "path" in response

def test_get_device():
    """Test the GET /devices endpoint."""
    global device_id

    response = requests.get(f"http://127.0.0.1:8888/devices/{device_id}")
    assert response.status_code == 200

    print(response.json())

    response = response.json()["data"]
    assert response["Name"] == data["Name"]
    assert response["Type"] == data["Type"]
    assert response["Location"] == data["Location"]
    assert response["Status"] == data["Status"]
    assert response["Description"] == data["Description"]
    assert "Device_Id" in response
    assert "Token" in response
    assert "Created_At" in response

def test_setup_device():
    """Test the GET /devices/:device_id/setup endpoint."""
    global device_id

    device_config = {
        "fields": {
            "temperature": "float64",
            "humidity": "int8"
        }
    }

    response = requests.post(f"http://127.0.0.1:8888/devices/{device_id}/setup", json=device_config)
    assert response.status_code == 200

    response = response.json()["data"]
    assert response["Device_Id"] == device_id
    assert isinstance(response["Fields"], dict)

def test_create_records_device():
    """Test the GET /devices/:device_id/records endpoint."""
    global device_id

    device_config = {
        "fields": {
                "temperature": 45.22,
                "humidity": 90
        }
    }

    response = requests.post(f"http://127.0.0.1:8888/devices/{device_id}/records", json=device_config)
    assert response.status_code == 200

    response = response.json()["data"]
    assert response["Device_Id"] == device_id
    assert isinstance(response["Fields"], dict)

def test_read_records_device():
    """Test the GET /devices/:device_id/records endpoint."""
    global device_id

    response = requests.get(f"http://127.0.0.1:8888/devices/{device_id}/records")
    assert response.status_code == 200

    response = response.json()["data"]
    assert isinstance(response, list)

def test_get_device_user():
    """Test the GET /devices/user endpoint."""
    global device_id

    response = requests.get(f"http://127.0.0.1:8888/devices/{device_id}/user")
    assert response.status_code == 200

    response = response.json()["data"]
    for user in response:
        assert isinstance(user, dict), "Each item in Users should be a dictionary"
        assert "User_Id" in user, "Each dictionary should have a 'User_Id' key"
        assert "Email" in user, "Each dictionary should have an 'Email' key"

def test_get_user_device():
    """Test the GET /users/{user_id}/devices endpoint."""
    global device_id
    owner = str(data["Owner"])

    response = requests.get(f"http://127.0.0.1:8888/users/{owner}/devices")
    assert response.status_code == 200

    response = response.json()["data"]
    for user in response:
        assert isinstance(user, dict), "Each item in Users should be a dictionary"
        assert "Device_Id" in user, "Each dictionary should have a 'Device_Id' key"
        assert "Name" in user, "Each dictionary should have an 'Name' key"
        assert "Token" in user, "Each dictionary should have an 'Token' key"

def test_delete_device():
    """Test the DELETE /devices endpoint."""
    global device_id

    response = requests.delete(f"http://127.0.0.1:8888/devices/{device_id}")
    owner = str(data["Owner"])

    requests.delete(f"http://localhost:8888/users/{owner}")

    assert response.status_code == 204
