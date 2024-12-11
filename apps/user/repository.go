package user

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

func (r repository) GetUserList(ctx context.Context) (users []User, err error) {
	query := `
		SELECT
			public_id, username, fullname, password_decoded, role
		FROM auth
		WHERE role='user'
	`

	err = r.db.SelectContext(ctx, &users, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}

	return

}

func (r repository) GetUserSaksiList(ctx context.Context) (users []User, err error) {
	query := `
		SELECT
			a.username,
			a.fullname,
			k.kecamatan_name,
			l.kelurahan_name,
			t.tps_name
		FROM tps t
			JOIN kelurahan l ON t.kelurahan_id = l.kelurahan_id
			JOIN kecamatan k ON l.kecamatan_id = k.kecamatan_id
			JOIN auth a ON t.user_id = a.public_id
		ORDER BY k.kecamatan_name ASC, l.kelurahan_name ASC, t.tps_name ASC
	`

	err = r.db.SelectContext(ctx, &users, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}

	return

}

func (r repository) EditTPSSaksi(ctx context.Context, model User, userId string) (err error) {

	var query string

	if model.Password != "" {
		query = `
		UPDATE auth SET fullname=$1, username=$2, password=$3, password_decoded=$4 WHERE public_id=$5
		`
		_, err = r.db.ExecContext(ctx, query, model.Fullname, model.Username, model.Password, model.PasswordDecoded, userId)
	} else {
		query = `
			UPDATE auth SET fullname=$1, username=$2 WHERE public_id=$3
			`
		_, err = r.db.ExecContext(ctx, query, model.Fullname, model.Username, userId)
	}

	if err != nil {
		return
	}

	return
}

func (r repository) GetDataForExportCSV(ctx context.Context) (users []User, err error) {
	query := `
	   SELECT
		k.kecamatan_name,
		l.kelurahan_name,
		t.tps_name,
		a.fullname,
		a.username,
		a.password_decoded
            FROM tps t
        JOIN kelurahan l ON t.kelurahan_id = l.kelurahan_id
        JOIN kecamatan k ON l.kecamatan_id = k.kecamatan_id
        JOIN auth a ON t.user_id = a.public_id
            ORDER BY k.kecamatan_name ASC, l.kelurahan_name ASC
	`

	err = r.db.SelectContext(ctx, &users, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}
	return

}

func (r repository) DeleteUserById(ctx context.Context, userId string) (err error) {

	query := `
		DELETE FROM auth WHERE public_id=$1 AND role='user'
	`

	if _, err = r.db.ExecContext(ctx, query, userId); err != nil {
		return
	}

	return nil
}
