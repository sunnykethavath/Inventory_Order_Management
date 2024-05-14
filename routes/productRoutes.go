package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunnykethavath/Inventory_Order_Management/controller"
)

func SetupProductRoutes(app *fiber.App) {

	app.Post("/admin/product", controller.CreateProduct)
	app.Get("/admin/products", controller.GetProducts)
	app.Get("/admin/product/:id", controller.GetProduct)
	app.Put("/admin/product/:id", controller.UpdateProduct)
	app.Delete("/admin/product/:id", controller.DeleteProduct)
}
