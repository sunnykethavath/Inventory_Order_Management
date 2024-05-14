package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunnykethavath/Inventory_Order_Management/database"
	"github.com/sunnykethavath/Inventory_Order_Management/routes"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	routes.SetupUserRoutes(app)

	routes.SetupProductRoutes(app)

	routes.SetupOrderRoutes(app)

	app.Listen(":3000")
}
