package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	PlayerID  uuid.UUID `json:"player_id"`
	Email     string    `json:"email"`
	Sub       string    `json:"sub"`
	IsBanned  bool      `json:"is_banned"`
	IsDeleted bool      `json:"is_deleted"`
}
