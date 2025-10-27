package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/domain"
)

func GetOrCreateUser(context context.Context, connection *pgx.Conn, email, sub string) (*domain.User, error) {
	query := `
        WITH inserted AS (
            INSERT INTO users (email, sub)
            VALUES ($1, $2)
            ON CONFLICT (sub) DO NOTHING
            RETURNING id, email, sub
        )
        SELECT id, email, sub FROM inserted
        UNION
        SELECT id, email, sub FROM users WHERE sub = $2;
    `

	var user domain.User
	err := connection.QueryRow(context, query, email, sub).Scan(&user.ID, &user.Email, &user.Sub)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
