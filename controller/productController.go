package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sunnykethavath/Inventory_Order_Management/database"
	"github.com/sunnykethavath/Inventory_Order_Management/model"
)

func CreateProduct(c *fiber.Ctx) error {
	var product model.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&product)

	return c.Status(200).JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
	products := []model.Product{}

	database.DB.Find(&products)

	return c.Status(200).JSON(products)
}

func UpdateProduct(c *fiber.Ctx) error {
	var requestProduct model.Product
	var product model.Product

	id, err1 := c.ParamsInt("id")

	if err1 != nil {
		c.Status(400).JSON("invalid id")
	}

	if err2 := c.BodyParser(&requestProduct); err2 != nil {
		return c.Status(400).JSON("cannot parse")
	}

	result := database.DB.Find(&product, "id = ?", id)

	if result.Error != nil {
		return c.Status(400).JSON("error updating product")
	}

	product.Name = requestProduct.Name
	product.Price = requestProduct.Price
	product.SerialNumber = requestProduct.SerialNumber

	database.DB.Save(&product)

	return c.Status(200).JSON(product)

}

func GetProduct(c *fiber.Ctx) error {
	var product model.Product

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("id error")
	}

	result := database.DB.Where("id = ?", id).First(&product)

	if result.Error != nil {
		c.Status(400).JSON("product does not exist")
	}

	if product.ID == 0 {
		c.Status(400).JSON("product does not exist")
	}

	return c.Status(200).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		c.Status(400).JSON("invalid id")
	}

	var product model.Product

	result := database.DB.Where("id = ?", id).First(&product)

	if result.Error != nil {
		c.Status(400).JSON("product does not exist")
	}

	result = database.DB.Delete(&product)

	if result.Error != nil {
		return c.Status(400).JSON("error deleting product")
	}

	jsonResponse := fmt.Sprintf("successfully deleted product with id %v", id)

	return c.Status(200).JSON(jsonResponse)
}
