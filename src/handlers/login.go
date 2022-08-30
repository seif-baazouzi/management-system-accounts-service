package handlers

import (
	"accounts-service/src/auth"
	"accounts-service/src/checks"
	"accounts-service/src/models"
	"accounts-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var body models.User

	// parse body
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check input
	errors := checks.CheckUser(&body)

	if errors != nil {
		return c.JSON(errors)
	}

	// check if not user is exist
	usersList, err := models.GetUsersByUsername(body.Username)

	if errors != nil {
		return utils.ServerError(c, err)
	}

	if len(usersList) == 0 {
		return c.JSON(fiber.Map{"username": "This username does not exist"})
	}

	user := usersList[0]

	// check password
	if !utils.ComparePasswords(body.Password, user.Password) {
		return c.JSON(fiber.Map{"username": "Wrong Password"})
	}

	// generate token
	token := auth.GenerateToken("uuid", user.UserID.String())

	return c.Status(201).JSON(fiber.Map{"message": "success", "token": token})
}
