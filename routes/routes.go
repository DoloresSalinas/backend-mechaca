package routes

import (
	"backend/handlers"
	"backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/usuarios")

	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)
 
	user := api.Group("/users", middleware.AuthMiddleware())
	user.Get("/", handlers.GetUsers) 
    user.Get("/:id", handlers.GetUser)
    user.Put("/:id", handlers.UpdateUser)
    user.Delete("/:id", handlers.DeleteUser)
}
