from testRestApi import runTests

import routes.signup as signup

if __name__ == "__main__":
    runTests([
        *signup.tests,
    ])
