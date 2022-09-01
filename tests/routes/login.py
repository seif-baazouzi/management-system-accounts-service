import utils
import config
from testRestApi import testRoute, Test, POST

def testEmptyFields():
    res = testRoute(POST, f"{config.server}/api/v1/login")
    return res.equals({ "username": "Must not be empty", "password": "Must not be empty" })

def testNotExistingUser():
    body = {
        "username": utils.randomString(10),
        "password": utils.randomString(10),
    }
    
    res = testRoute(POST, f"{config.server}/api/v1/login", body=body)

    return res.equals({ "username": "This username does not exist"})

def testWrongPassword():
    username, password, token = utils.signup()

    body = {
        "username": username,
        "password": utils.randomString(10),
    }
    res = testRoute(POST, f"{config.server}/api/v1/login", body=body)
    
    return res.equals({ "password": "Wrong Password"})

def testLogin():
    body = {
        "username": utils.randomString(10),
        "password": utils.randomString(10),
    }

    testRoute(POST, f"{config.server}/api/v1/signup", body=body)
    res = testRoute(POST, f"{config.server}/api/v1/login", body=body)

    return res.hasKeys("message", "token") and res.equals({ "message": "success" })

tests = [
    Test("Login Empty Fields", testEmptyFields),
    Test("Login Not Existing User", testNotExistingUser),
    Test("Login Wrong Password", testWrongPassword),
    Test("Login", testLogin),
]
