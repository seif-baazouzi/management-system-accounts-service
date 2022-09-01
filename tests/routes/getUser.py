import utils
import config
from testRestApi import testRoute, Test, GET

def testToken():
    res = testRoute(GET, f"{config.server}/api/v1/user")
    return res.equals({ "message": "invalid-token" })

def testGetUser():
    username, password, token = utils.signup()
    res = testRoute(GET, f"{config.server}/api/v1/user", headers={ "X-Token": token })
        
    return res.equals({ "username": username })

tests = [
    Test("Get User Invalid Token", testToken),
    Test("Get User", testGetUser),
]
