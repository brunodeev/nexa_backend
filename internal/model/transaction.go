package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	CategoryID  uuid.UUID `json:"category_id" db:"category_id"`
	Amount      float32   `json:"amount" db:"amount"`
	Date        time.Time `json:"date" db:"date"`
	Description string    `json:"description" db:"description"`
	Type        string    `json:"type" db:"type"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
