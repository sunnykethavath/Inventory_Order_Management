package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunnykethavath/Inventory_Order_Management/database"
	"github.com/sunnykethavath/Inventory_Order_Management/model"
)

func CreateOrder(c *fiber.Ctx) error {
	var order model.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user model.User

	if err := database.DB.Where("id = ?", &order.UserRefer).First(&user); err.Error != nil {
		return c.Status(400).JSON("user does not exist")
	}

	var product model.Product

	if err := database.DB.Where("id = ?", &order.ProductRefer).First(&product); err.Error != nil {
		return c.Status(400).JSON("product doesn't exist")
	}

	order.User = user
	order.Product = product

	database.DB.Create(&order)

	return c.Status(200).JSON(order)
}

func GetOrders(c *fiber.Ctx) error {
	var orders []model.Order

	database.DB.Find(&orders)

	responseOrders := []model.Order{}
	for _, order := range orders {
		var user model.User
		var product model.Product
		database.DB.Where("id = ?", order.UserRefer).First(&user)
		database.DB.Where("id = ?", order.ProductRefer).First(&product)
		order.User = user
		order.Product = product
		responseOrders = append(responseOrders, order)
	}
	return c.Status(200).JSON(responseOrders)
}

func GetOrder(c *fiber.Ctx) error {
	var order model.Order

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := database.DB.Where("id = ?", id).First(&order); err.Error != nil {
		return c.Status(400).JSON("order does not exist")
	}

	var user model.User
	var product model.Product

	database.DB.Where("id = ?", order.UserRefer).First(&user)
	database.DB.Where("id = ?", order.ProductRefer).First(&product)

	order.User = user
	order.Product = product

	return c.Status(200).JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	var order model.Order

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	database.DB.Where("id = ?", id).First(&order)

	if order.ID == 0 {
		return c.Status(400).JSON("user does not exist")
	}

	result := database.DB.Delete(&order)

	if result.Error != nil {
		return c.Status(400).JSON("error deleting user")
	}

	var user model.User
	var product model.Product

	database.DB.Where("id = ?", order.UserRefer).First(&user)
	database.DB.Where("id = ?", order.ProductRefer).First(&product)

	order.User = user
	order.Product = product

	return c.Status(200).JSON(order)
}
