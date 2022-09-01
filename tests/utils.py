import string
import random

def randomString(length):
    return ''.join(random.choices(string.ascii_letters, k=length))
