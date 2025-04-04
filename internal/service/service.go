package service

import (
	"nexa/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func StartServer(conn *pgx.Conn) {
	app := fiber.New()
	uh := handler.NewUserHandler(conn)
	api := app.Group("/api/users")

	api.Post("/", uh.CreateUser)
	// api.Get("/", handler.GetUsers)
	// api.Get("/:id", handler.GetUserByID)
	// api.Put("/:id", handler.UpdateUser)
	// api.Delete("/:id", handler.DeleteUser)

	app.Listen(":8080")
}
