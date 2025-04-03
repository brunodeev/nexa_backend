package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"nexa/internal/handler"
	"nexa/internal/database"
)

func StartServer() {
	
	database.ConnectDB()

	app := fiber.New()

	api := app.Group("/api/users")
	api.Post("/", handler.CreateUser)
	api.Get("/", handler.GetUsers)
	api.Get("/:id", handler.GetUserByID)
	api.Put("/:id", handler.UpdateUser)
	api.Delete("/:id", handler.DeleteUser)

	log.Println("Servidor rodando na porta 8080")
	app.Listen(":8080")
}
