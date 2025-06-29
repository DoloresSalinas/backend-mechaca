package main

import (
	"backend/config"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inicializar conexión a DB
	config.InitDB()
	defer config.Conn.Close(nil)

	app := fiber.New()

	// Configurar rutas
	routes.SetupUserRoutes(app)

	app.Listen(":3000")
}
