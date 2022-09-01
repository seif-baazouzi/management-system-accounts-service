import string
import random

import config
from testRestApi import testRoute, POST

def randomString(length):
    return ''.join(random.choices(string.ascii_letters, k=length))

def signup():
    username = randomString(10)
    password = randomString(10)

    body = { "username": username, "password": password }
    res = testRoute(POST, f"{config.server}/api/v1/signup", body=body)
    token = False if not res.hasKeys("token") else res.body["token"]

    return username, password, token

def login(username, password):
    body = { "username": username, "password": password }
    res = testRoute(POST, f"{config.server}/api/v1/login", body=body)
    token = False if not res.hasKeys("token") else res.body["token"]

    return token

