import pytest
import requests

user_id = None
data = {
    "Name":"Alpha Zero One",
    "Email":"alpha01@mail.com",
    "Password":"password"
}

def test_create_user():
    """Test the POST /users endpoint."""
    global user_id

    response = requests.post("http://127.0.0.1:8888/users", json=data)
    assert response.status_code == 201 

    response = response.json()["data"]
    assert response["Name"] == data["Name"]
    assert response["Email"] == data["Email"]
    assert "Password_Hash" in response
    assert "User_Id" in response

    user_id = response["User_Id"]

def test_get_user():
    """Test the GET /users/info endpoint."""
    global user_id

    response = requests.get(f"http://127.0.0.1:8888/users/{user_id}/info")
    assert response.status_code == 200

    response = response.json()["data"]
    assert response["Name"] == data["Name"]
    assert response["Email"] == data["Email"]
    assert "User_Id" in response
    assert "Created_At" in response

def test_login_user():
    """Test the POST /auth/token endpoint."""
    global user_id

    authData = {
        "Email": data["Email"],
        "Password": data["Password"]
    }

    response = requests.post("http://127.0.0.1:8888/auth/token", json=authData)
    assert response.status_code == 200

    response = response.json()["data"]
    assert response["Name"] == data["Name"]
    assert response["Email"] == data["Email"]
    assert response["User_Id"] == user_id
    assert "Token" in response

def test_create_user_invalid_data():
    """Test the POST /users endpoint with invalid data."""

    invalidData = {
        "email": "invalid_user@example.com" 
    }

    response = requests.post("http://127.0.0.1:8888/users", json=invalidData)
    assert response.status_code == 400

    response = response.json()["errors"]
    assert "timestamp" in response
    assert "status" in response
    assert "error" in response
    assert "message" in response
    assert "path" in response

def test_login_user_invalid_data():
    """Test the POST /auth/token endpoint with invalid data."""

    invalidData = {
        "email": "invalid_user@example.com" 
    }

    response = requests.post("http://127.0.0.1:8888/auth/token", json=invalidData)
    assert response.status_code == 400

    response = response.json()["errors"]
    assert "timestamp" in response
    assert "status" in response
    assert "error" in response
    assert "message" in response
    assert "path" in response

def test_delete_user_info():
    """Test the DELETE /users endpoint."""
    global user_id

    response = requests.delete(f"http://localhost:8888/users/{user_id}")

    assert response.status_code == 204
