package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Email string    `json:"email" db:"email"`
	Sub   string    `json:"sub" db:"sub"`
}
