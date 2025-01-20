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
    assert response.json()["Name"] == data["Name"]
    assert response.json()["Email"] == data["Email"]
    assert "Password_Hash" in response.json()
    assert "User_Id" in response.json()

    user_id = response.json()["User_Id"]

def test_get_user():
    """Test the GET /users/info endpoint."""
    global user_id

    response = requests.get(f"http://127.0.0.1:8888/users/info?user_id={user_id}")

    assert response.status_code == 200
    assert response.json()["Name"] == data["Name"]
    assert response.json()["Email"] == data["Email"]
    assert "User_Id" in response.json()
    assert "Created_At" in response.json()

def test_login_user():
    """Test the POST /auth/token endpoint."""
    global user_id

    authData = {
        "Email": data["Email"],
        "Password": data["Password"]
    }

    response = requests.post("http://127.0.0.1:8888/auth/token", json=authData)

    assert response.status_code == 200
    assert response.json()["Name"] == data["Name"]
    assert response.json()["Email"] == data["Email"]
    assert response.json()["User_Id"] == user_id
    assert "Token" in response.json()

def test_create_user_invalid_data():
    """Test the POST /users endpoint with invalid data."""

    invalidData = {
        "email": "invalid_user@example.com" 
    }

    response = requests.post("http://127.0.0.1:8888/users", json=invalidData)

    assert response.status_code == 400
    assert "timestamp" in response.json()
    assert "status" in response.json()
    assert "error" in response.json()
    assert "message" in response.json()
    assert "path" in response.json()

def test_login_user_invalid_data():
    """Test the POST /auth/token endpoint with invalid data."""

    invalidData = {
        "email": "invalid_user@example.com" 
    }

    response = requests.post("http://127.0.0.1:8888/auth/token", json=invalidData)

    assert response.status_code == 400
    assert "timestamp" in response.json()
    assert "status" in response.json()
    assert "error" in response.json()
    assert "message" in response.json()
    assert "path" in response.json()

def test_delete_user_info():
    """Test the GET /users endpoint with valid data."""
    global user_id

    response = requests.delete(f"http://localhost:8888/users?user_id={user_id}")

    assert response.status_code == 204
