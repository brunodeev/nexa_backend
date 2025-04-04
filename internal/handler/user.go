package handler

import (
	"context"
	"time"

	"nexa/internal/database"
	"nexa/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	query := `INSERT INTO users (id, name, email, password, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao criar usuário"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	rows, err := database.DB.Query(context.Background(), "SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao buscar usuários"})
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao processar usuários"})
		}
		users = append(users, user)
	}

	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var user model.User
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
	err := database.DB.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	user.UpdatedAt = time.Now()

	query := "UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4"
	_, err := database.DB.Exec(context.Background(), query, user.Name, user.Email, user.UpdatedAt, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao atualizar usuário"})
	}

	return c.JSON(fiber.Map{"message": "Usuário atualizado com sucesso"})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	query := "DELETE FROM users WHERE id = $1"
	_, err := database.DB.Exec(context.Background(), query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao excluir usuário"})
	}

	return c.JSON(fiber.Map{"message": "Usuário excluído com sucesso"})
}
