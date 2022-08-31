package checks

import (
	"accounts-service/src/models"

	"github.com/gofiber/fiber/v2"
)

func CheckUpdatePassword(user *models.UpdatePassword) fiber.Map {
	errors := make(fiber.Map)

	if user.OldPassword == "" {
		errors["oldPassword"] = "Must not be empty"
	}

	if user.NewPassword == "" {
		errors["newPassword"] = "Must not be empty"
	} else if user.NewPassword == user.OldPassword {
		errors["newPassword"] = "Must not be the same as the old one"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}
