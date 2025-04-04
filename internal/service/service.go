package service

import (
	"nexa/internal/handler"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func StartServer(conn *pgx.Conn) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()
	uh := handler.NewUserHandler(conn)
	api := app.Group("/api/users")

	api.Post("/", uh.CreateUser)
	api.Get("/", uh.GetUsers)
	// api.Get("/:id", handler.GetUserByID)
	// api.Put("/:id", handler.UpdateUser)
	// api.Delete("/:id", handler.DeleteUser)

	app.Listen(":" + port)
}
