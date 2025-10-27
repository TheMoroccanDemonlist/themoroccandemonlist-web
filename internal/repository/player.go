package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/domain"
)

func GetOrCreatePlayer(context context.Context, connection *pgx.Conn, userID uuid.UUID) (*domain.Player, error) {
	query := `
        WITH inserted AS (
            INSERT INTO players (user_id)
            VALUES ($1)
            ON CONFLICT (user_id) DO NOTHING
            RETURNING id, user_id
        )
        SELECT id, user_id FROM inserted
        UNION
        SELECT id, user_id FROM players WHERE user_id = $1;
    `

	var player domain.Player
	err := connection.QueryRow(context, query, userID).Scan(&player.ID, &player.UserID)
	if err != nil {
		return nil, err
	}
	return &player, nil
}
