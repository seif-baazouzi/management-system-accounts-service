package handlers

import (
	"accounts-service/src/checks"
	"accounts-service/src/models"
	"accounts-service/src/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UpdatePassword(c *fiber.Ctx) error {
	var body models.UpdatePassword

	uuid := fmt.Sprintf("%s", c.Locals("uuid"))

	// parse body
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check input
	errors := checks.CheckUpdatePassword(&body)

	if errors != nil {
		return c.JSON(errors)
	}

	// check password
	user, err := models.GetUserByUserID(uuid)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !utils.ComparePasswords(body.OldPassword, user.Password) {
		return c.JSON(fiber.Map{"password": "Wrong Password"})
	}

	// update password
	models.UpdateUser(&models.User{
		UserID:   user.UserID,
		Username: user.Username,
		Password: body.NewPassword,
	})

	return c.Status(201).JSON(fiber.Map{"message": "success"})
}
