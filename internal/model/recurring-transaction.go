package model

import (
	"time"

	"github.com/google/uuid"
)

type RecurringTransaction struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	CategoryID  uuid.UUID `json:"category_id" db:"category_id"`
	Amount      float32   `json:"amount" db:"amount"`
	Description string    `json:"description" db:"description"`
	Type        string    `json:"type" db:"type"`
	StartDate   time.Time `json:"start_date" db:"start_date"`
	EndDate     time.Time `json:"end_date" db:"end_date"`
	NextRunDate time.Time `json:"next_run_date" db:"next_run_date"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
