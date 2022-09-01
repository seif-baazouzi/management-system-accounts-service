package handlers

import (
	"accounts-service/src/models"
	"accounts-service/src/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	uuid := fmt.Sprintf("%s", c.Locals("uuid"))

	// get user
	user, err := models.GetUserByUserID(uuid)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"userID": user.UserID, "username": user.Username})
}
