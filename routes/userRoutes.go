package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunnykethavath/Inventory_Order_Management/controller"
)

func SetupUserRoutes(app *fiber.App) {

	app.Post("/user/profile", controller.CreateUser)
	app.Get("/user/profiles", controller.GetUsers)
	app.Delete("/user/profile/:id", controller.DeleteUser)
}
