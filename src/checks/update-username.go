package checks

import (
	"accounts-service/src/models"

	"github.com/gofiber/fiber/v2"
)

func CheckUpdateUsername(user *models.UpdateUsername) fiber.Map {
	errors := make(fiber.Map)

	if user.NewUsername == "" {
		errors["newUsername"] = "Must not be empty"
	}

	if user.Password == "" {
		errors["password"] = "Must not be empty"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}
