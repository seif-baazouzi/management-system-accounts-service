import utils
import config
from testRestApi import testRoute, Test, PUT

def testToken():
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/password")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    username, password, token = utils.signup()    
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/password", headers={ "X-Token": token })
    
    return res.equals({ "oldPassword": "Must not be empty", "newPassword": "Must not be empty" })

def testSamePassword():
    username, password, token = utils.signup()

    body = {
        "oldPassword": password,
        "newPassword": password,
    }
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/password", headers={ "X-Token": token }, body=body)
    return res.equals({ "newPassword": "Must not be the same as the old one" })

def testInvalidPassword():
    username, password, token = utils.signup()

    body = {
        "oldPassword": utils.randomString(10),
        "newPassword": utils.randomString(10),
    }
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/password", headers={ "X-Token": token }, body=body)
    return res.equals({ "oldPassword": "Wrong Password" })

def testUpdatePassword():
    username, password, token = utils.signup()
    newPassword = utils.randomString(10)

    body = {
        "oldPassword": password,
        "newPassword": newPassword,
    }
    res = testRoute(PUT, f"{config.server}/api/v1/settings/update/password", headers={ "X-Token": token }, body=body)
    if not res.equals({ "message": "success" }):
        return False

    token = utils.login(username, newPassword)
    return token != False

tests = [
    Test("Update Password Invalid Token", testToken),
    Test("Update Password Empty Fields", testEmptyFields),
    Test("Update Password Same Password", testSamePassword),
    Test("Update Password Invalid Password", testInvalidPassword),
    Test("Update Password", testUpdatePassword),
]
