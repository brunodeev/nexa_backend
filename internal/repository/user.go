package repository

import (
	"context"
	"fmt"
	"nexa/internal/model"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	Conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (ur *UserRepository) InsertUser(user *model.User) error {
	query := `INSERT INTO "user" (id, name, email, password, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := ur.Conn.Exec(
		context.Background(),
		query,
		user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	rows, err := ur.Conn.Query(context.Background(), `SELECT id, name, email, created_at, updated_at FROM "user"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
