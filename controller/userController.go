package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sunnykethavath/Inventory_Order_Management/database"
	"github.com/sunnykethavath/Inventory_Order_Management/model"
)

func CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&user)

	return c.Status(200).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	users := []model.User{}

	database.DB.Find(&users)

	return c.Status(200).JSON(users)
}

func DeleteUser(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		c.Status(400).JSON("invalid id")
	}

	var user model.User

	database.DB.Where("id = ?", id).First(&user)

	if user.ID == 0 {
		return c.Status(400).JSON("user does not exist")
	}

	result := database.DB.Delete(&user)

	if result.Error != nil {
		return c.Status(400).JSON("error deleting user")
	}

	jsonResponse := fmt.Sprintf("successfully deleted user with id %v", id)

	return c.Status(200).JSON(jsonResponse)

}
