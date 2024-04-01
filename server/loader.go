package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	SetupRoutes(app)
}

func StartServer(port int) {
	app := fiber.New()
	LoadRoutes(app)
	app.Listen(fmt.Sprintf(":%d", port))
	fmt.Printf("Server started on port %d\n - http://localhost:%d", port, port)
}
