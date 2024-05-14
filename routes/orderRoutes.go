package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunnykethavath/Inventory_Order_Management/controller"
)

func SetupOrderRoutes(app *fiber.App) {
	app.Post("/user/order", controller.CreateOrder)
	app.Get("/user/orders", controller.GetOrders)
	app.Get("/user/order/:id", controller.GetOrder)
	app.Delete("/user/order/:id", controller.DeleteOrder)
}
