package handlers

import (
	"accounts-service/src/models"
	"accounts-service/src/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	var body models.DeleteUser

	uuid := fmt.Sprintf("%s", c.Locals("uuid"))

	// parse body
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check input
	if body.Password == "" {
		return c.JSON(fiber.Map{"password": "Must not be empty"})
	}

	// check password
	user, err := models.GetUserByUserID(uuid)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !utils.ComparePasswords(body.Password, user.Password) {
		return c.JSON(fiber.Map{"password": "Wrong Password"})
	}

	// delete user
	models.DeleteUserByUserID(uuid)

	return c.Status(201).JSON(fiber.Map{"message": "success"})
}
