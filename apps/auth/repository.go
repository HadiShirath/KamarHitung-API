package auth

import (
	"context"
	"database/sql"
	"kamar-hitung/infra/response"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateAuth(ctx context.Context, model AuthEntity) (err error) {
	query := `
		INSERT INTO auth (
			username, password, password_decoded, role, fullname, created_at, updated_at, public_id
		) VALUES (
			:username, :password, :password_decoded, :role, :fullname, :created_at, :updated_at, :public_id
		)
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)

	return
}

func (r repository) GetAuthByUsername(ctx context.Context, username string) (model AuthEntity, err error) {
	query := `
		SELECT
			public_id, username, password, fullname, role, created_at, updated_at 
		FROM auth
		WHERE username=$1
	`

	err = r.db.GetContext(ctx, &model, query, username)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}
