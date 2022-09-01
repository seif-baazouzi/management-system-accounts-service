from testRestApi import runTests

import routes.login as login
import routes.signup as signup
import routes.updateUsername as updateUsername
import routes.updatePassword as updatePassword
import routes.deleteUser as deleteUser

if __name__ == "__main__":
    runTests([
        *signup.tests,
        *login.tests,
        *updateUsername.tests,
        *updatePassword.tests,
        *deleteUser.tests,
    ])
