import utils
import config
from testRestApi import testRoute, Test, POST, PUT

def testToken():
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    body = {
        "username": utils.randomString(10),
        "password": utils.randomString(10),
    }
    res = testRoute(POST, f"{config.server}/api/v1/signup", body=body)
    token = res.body["token"]

    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username", headers={ "X-Token": token })
    return res.equals({ "newUsername": "Must not be empty", "password": "Must not be empty" })

def testExistingUsername():
    user1 = {
        "username": utils.randomString(10),
        "password": utils.randomString(10),
    }
    testRoute(POST, f"{config.server}/api/v1/signup", body=user1)
    
    signupBody = {
        "username": utils.randomString(10),
        "password": utils.randomString(10),
    }
    res = testRoute(POST, f"{config.server}/api/v1/signup", body=signupBody)
    token = res.body["token"]

    body = {
        "newUsername": user1["username"],
        "password": signupBody["password"],
    }
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username", headers={ "X-Token": token }, body=body)
    return res.equals({ "newUsername": "This username is already exist" })

def testInvalidPassword():
    signupBody = {
        "username": utils.randomString(10),
        "password": utils.randomString(10),
    }
    res = testRoute(POST, f"{config.server}/api/v1/signup", body=signupBody)
    token = res.body["token"]

    body = {
        "newUsername": utils.randomString(10),
        "password": utils.randomString(10),
    }
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username", headers={ "X-Token": token }, body=body)
    return res.equals({ "password": "Wrong Password" })

def testUpdateUsername():
    signupBody = {
        "username": utils.randomString(10),
        "password": utils.randomString(10),
    }
    res = testRoute(POST, f"{config.server}/api/v1/signup", body=signupBody)
    token = res.body["token"]

    body = {
        "newUsername": utils.randomString(10),
        "password": signupBody["password"],
    }
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username", headers={ "X-Token": token }, body=body)
    if not res.equals({ "message": "success" }):
        return False

    loginBody = {
        "username": body["newUsername"],
        "password": signupBody["password"],
    }
    res = testRoute(POST, f"{config.server}/api/v1/login", body=loginBody)
    return res.hasKeys("message", "token") and res.equals({ "message": "success" })

tests = [
    Test("Update Username Invalid Token", testToken),
    Test("Update Username Empty Fields", testEmptyFields),
    Test("Update Username Existing Username", testExistingUsername),
    Test("Update Username Invalid Password", testInvalidPassword),
    Test("Update Username", testUpdateUsername),
]
