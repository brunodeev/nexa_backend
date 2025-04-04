package handler

import (
	"nexa/internal/factory"
	"nexa/internal/model"
	"nexa/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type UserHandler struct {
	Conn           *pgx.Conn
	UserRepository repository.UserRepository
	UserFactory    factory.UserFactory
}

func NewUserHandler(conn *pgx.Conn) *UserHandler {
	return &UserHandler{
		Conn:           conn,
		UserRepository: *repository.NewUserRepository(conn),
		UserFactory:    *factory.NewUserFactory(),
	}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	body := new(model.User)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	user := uh.UserFactory.CreateUser(*body)

	err := uh.UserRepository.InsertUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(user)
}

// func (uh *UserHandler) GetUsers(c *fiber.Ctx) error {
// 	rows, err := database.DB.Query(context.Background(), "SELECT id, name, email, created_at, updated_at FROM users")
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao buscar usuários"})
// 	}
// 	defer rows.Close()

// 	var users []model.User
// 	for rows.Next() {
// 		var user model.User
// 		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao processar usuários"})
// 		}
// 		users = append(users, user)
// 	}

// 	return c.JSON(users)
// }

// func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	var user model.User
// 	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
// 	err := database.DB.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Usuário não encontrado"})
// 	}

// 	return c.JSON(user)
// }

// func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var user model.User

// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos"})
// 	}

// 	user.UpdatedAt = time.Now()

// 	query := "UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4"
// 	_, err := database.DB.Exec(context.Background(), query, user.Name, user.Email, user.UpdatedAt, id)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao atualizar usuário"})
// 	}

// 	return c.JSON(fiber.Map{"message": "Usuário atualizado com sucesso"})
// }

// func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	query := "DELETE FROM users WHERE id = $1"
// 	_, err := database.DB.Exec(context.Background(), query, id)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao excluir usuário"})
// 	}

// 	return c.JSON(fiber.Map{"message": "Usuário excluído com sucesso"})
// }
