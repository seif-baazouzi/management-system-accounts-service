package handlers

import (
	"accounts-service/src/auth"
	"accounts-service/src/checks"
	"accounts-service/src/models"
	"accounts-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
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

	// check if user is exist
	usersList, err := models.GetUsersByUsername(body.Username)

	if errors != nil {
		return utils.ServerError(c, err)
	}

	if len(usersList) != 0 {
		return c.JSON(fiber.Map{"username": "This username is already exist"})
	}

	// create user
	userID, err := models.CreateUser(&body)

	if errors != nil {
		return utils.ServerError(c, err)
	}

	// generate token
	token := auth.GenerateToken("uuid", userID)

	return c.Status(201).JSON(fiber.Map{"message": "success", "token": token})
}
