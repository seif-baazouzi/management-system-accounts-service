package handlers

import (
	"accounts-service/src/checks"
	"accounts-service/src/models"
	"accounts-service/src/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UpdateUsername(c *fiber.Ctx) error {
	var body models.UpdateUsername

	uuid := fmt.Sprintf("%s", c.Locals("uuid"))

	// parse body
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check input
	errors := checks.CheckUpdateUsername(&body)

	if errors != nil {
		return c.JSON(errors)
	}

	// check if not user is exist
	usersList, err := models.GetUsersByUsername(body.NewUsername)

	if errors != nil {
		return utils.ServerError(c, err)
	}

	if len(usersList) != 0 {
		return c.JSON(fiber.Map{"newUsername": "This username is already exist"})
	}

	// check password
	user, err := models.GetUserByUserID(uuid)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !utils.ComparePasswords(body.Password, user.Password) {
		return c.JSON(fiber.Map{"password": "Wrong Password"})
	}

	// update username
	models.UpdateUser(&models.User{
		UserID:   user.UserID,
		Username: body.NewUsername,
		Password: body.Password,
	})

	return c.JSON(fiber.Map{"message": "success"})
}
