import utils
import config
from testRestApi import testRoute, Test, PUT

def testToken():
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    username, password, token = utils.signup()    
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username", headers={ "X-Token": token })
    
    return res.equals({ "newUsername": "Must not be empty", "password": "Must not be empty" })

def testExistingUsername():
    user1Username, user1Password, _ = utils.signup()
    username, password, token = utils.signup()
    
    body = {
        "newUsername": user1Username,
        "password": password,
    }
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username", headers={ "X-Token": token }, body=body)
    return res.equals({ "newUsername": "This username is already exist" })

def testInvalidPassword():
    username, password, token = utils.signup()

    body = {
        "newUsername": utils.randomString(10),
        "password": utils.randomString(10),
    }
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username", headers={ "X-Token": token }, body=body)

    return res.equals({ "password": "Wrong Password" })

def testUpdateUsername():
    username, password, token = utils.signup()

    body = {
        "newUsername": utils.randomString(10),
        "password": password,
    }
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/username", headers={ "X-Token": token }, body=body)
    if not res.equals({ "message": "success" }):
        return False

    username = body["newUsername"]

    token = utils.login(username, password)
    return token != False

tests = [
    Test("Update Username Invalid Token", testToken),
    Test("Update Username Empty Fields", testEmptyFields),
    Test("Update Username Existing Username", testExistingUsername),
    Test("Update Username Invalid Password", testInvalidPassword),
    Test("Update Username", testUpdateUsername),
]
