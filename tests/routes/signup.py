import utils
import config
from testRestApi import testRoute, Test, POST

def testEmptyFields():
    res = testRoute(POST, f"{config.server}/api/v1/signup")
    return res.equals({ "username": "Must not be empty", "password": "Must not be empty" })

def testExistingUser():
    body = {
        "username": utils.randomString(10),
        "password": utils.randomString(10),
    }
    
    testRoute(POST, f"{config.server}/api/v1/signup", body=body)
    res = testRoute(POST, f"{config.server}/api/v1/signup", body=body)

    return res.equals({ "username": "This username is already exist"})

def testSignup():
    res = testRoute(POST, f"{config.server}/api/v1/signup", body={
        "username": utils.randomString(10),
        "password": utils.randomString(10),
    })

    return res.hasKeys("message", "token") and res.equals({ "message": "success" })

tests = [
    Test("Signup Empty Fields", testEmptyFields),
    Test("Signup Existing User", testExistingUser),
    Test("Signup", testSignup),
]
