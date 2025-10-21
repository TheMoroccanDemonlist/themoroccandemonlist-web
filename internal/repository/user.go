package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/domain"
)

func GetOrCreateUser(context context.Context, connection *pgx.Conn, email, sub string) (*domain.User, error) {
	query := `
        INSERT INTO users (email, sub)
        VALUES ($1, $2)
        ON CONFLICT (sub) DO UPDATE
        SET email = EXCLUDED.email
        RETURNING id, email, sub;
    `

	var user domain.User
	err := connection.QueryRow(context, query, email, sub).Scan(&user.ID, &user.Email, &user.Sub)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
