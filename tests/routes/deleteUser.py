import utils
import config
from testRestApi import testRoute, Test, DELETE

def testToken():
    res = testRoute(DELETE, f"{config.server}/api/v1/settings/delete-user")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    username, password, token = utils.signup()    
    res = testRoute(DELETE, f"{config.server}/api/v1/settings/delete-user", headers={ "X-Token": token })
    
    return res.equals({ "password": "Must not be empty" })

def testInvalidPassword():
    username, password, token = utils.signup()

    body = { "password": utils.randomString(10) }
    res = testRoute(DELETE, f"{config.server}/api/v1/settings/delete-user", headers={ "X-Token": token }, body=body)
    return res.equals({ "password": "Wrong Password" })

def testDeleteUser():
    username, password, token = utils.signup()
    newPassword = utils.randomString(10)

    body = { "password": password }
    res = testRoute(DELETE, f"{config.server}/api/v1/settings/delete-user", headers={ "X-Token": token }, body=body)
    if not res.equals({ "message": "success" }):
        return False

    token = utils.login(username, password)
    return token == False

tests = [
    Test("Delete User Invalid Token", testToken),
    Test("Delete User Empty Fields", testEmptyFields),
    Test("Delete User Invalid Password", testInvalidPassword),
    Test("Delete User", testDeleteUser),
]
