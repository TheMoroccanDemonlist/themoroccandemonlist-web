package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Player struct {
	ID               uuid.UUID       `json:"id"`
	UserID           uuid.UUID       `json:"user_id"`
	Username         string          `json:"username"`
	Avatar           string          `json:"avatar"`
	ClassicPoints    decimal.Decimal `json:"classic_points"`
	PlatformerPoints decimal.Decimal `json:"platformer_points"`
	Discord          string          `json:"discord"`
	YouTube          string          `json:"youtube"`
	Twitter          string          `json:"twitter"`
	Twitch           string          `json:"twitch"`
	IsFlagged        bool            `json:"is_flagged"`
}
