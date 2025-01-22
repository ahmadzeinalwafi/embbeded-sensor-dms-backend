import pytest
import requests

device_id = None
data = {
	"Name": "MeasurementIoT",
	"Type": "Sensor",
	"Location": "Field A",
	"Status": "Active",
	"Description": "This measure humidity and temperature",
	"Owner": "xfHf39PM"
}

def test_create_device():
    """Test the POST /devices endpoint."""
    global device_id

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

    response = response.json()["data"]
    assert response["Name"] == data["Name"]
    assert response["Type"] == data["Type"]
    assert response["Location"] == data["Location"]
    assert response["Status"] == data["Status"]
    assert response["Description"] == data["Description"]
    assert "Device_Id" in response
    assert "Token" in response
    assert "Created_At" in response

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

def test_delete_device():
    """Test the DELETE /devices endpoint."""
    global device_id

    response = requests.delete(f"http://127.0.0.1:8888/devices/{device_id}")

    assert response.status_code == 204
