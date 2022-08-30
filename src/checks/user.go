package checks

import (
	"accounts-service/src/models"

	"github.com/gofiber/fiber/v2"
)

func CheckUser(user *models.User) fiber.Map {
	errors := make(fiber.Map)

	if user.Username == "" {
		errors["username"] = "Must not be empty"
	}

	if user.Password == "" {
		errors["password"] = "Must not be empty"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}
