package kecamatan

import (
	"context"
	"database/sql"
	"fmt"
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

func (r repository) GetVoterKecamatan(ctx context.Context, codeKecamatan string) (kecamatan Kecamatan, err error) {
	query := `
		SELECT
    k.kecamatan_name,
    SUM(t.paslon1) AS paslon1,
    SUM(t.paslon2) AS paslon2,
    SUM(t.paslon3) AS paslon3,
    SUM(t.paslon4) AS paslon4,
    SUM(t.suara_tidak_sah) AS suara_tidak_sah
		FROM
			tps t
		JOIN
			kecamatan k ON t.kecamatan_id = k.kecamatan_id
		WHERE
			k.code = $1
		GROUP BY
			k.kecamatan_name;
	`
	err = r.db.GetContext(ctx, &kecamatan, query, codeKecamatan)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) GetAllVoter(ctx context.Context) (kecamatan Kecamatan, err error) {
	query := `
		SELECT
    SUM(suara_sah + suara_tidak_Sah) AS total_suara,
    ROUND(
        (COUNT(CASE WHEN suara_sah IS NOT NULL AND suara_sah > 0 THEN 1 END) * 100.0 / COUNT(*))
    ) AS persentase
		FROM tps;
	`
	err = r.db.GetContext(ctx, &kecamatan, query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) GetListKecamatan(ctx context.Context) (kecamatans []Kecamatan, err error) {
	query := `
		SELECT
			c.kecamatan_name,
			COALESCE(SUM(t.paslon1), 0) AS paslon1,
			COALESCE(SUM(t.paslon2), 0) AS paslon2,
			COALESCE(SUM(t.paslon3), 0) AS paslon3,
			COALESCE(SUM(t.paslon4), 0) AS paslon4,
			COALESCE(SUM(t.suara_sah), 0) AS suara_sah,
			COALESCE(SUM(t.suara_tidak_sah), 0) AS suara_tidak_sah,
			c.total_voters,
			c.total_tps,
			COUNT(CASE WHEN t.suara_sah > 0 THEN 1 END) AS sudah,
			COUNT(CASE WHEN t.suara_sah = 0 THEN 1 END) AS belum,
			ROUND(
				(COALESCE(SUM(t.suara_sah), 0) + COALESCE(SUM(t.suara_tidak_sah), 0)) * 100.0 / COALESCE(c.total_voters, 1)
			) AS pp,
			c.code
		FROM
			kecamatan c
		LEFT JOIN
			kelurahan k ON c.kecamatan_id = k.kecamatan_id
		LEFT JOIN
			tps t ON k.kelurahan_id = t.kelurahan_id
		GROUP BY
			c.kecamatan_id, c.kecamatan_name, c.total_voters, c.total_tps
		ORDER BY
			c.kecamatan_name;
	`

	err = r.db.SelectContext(ctx, &kecamatans, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}
	return

}

func (r repository) GetListKelurahanFromKecamatan(ctx context.Context, codeKecamatan string) (kelurahans []Kelurahan, err error) {
	query := `
	    SELECT
			c.kecamatan_name,
			k.kelurahan_name,
			COALESCE(SUM(t.paslon1), 0) AS paslon1,
			COALESCE(SUM(t.paslon2), 0) AS paslon2,
			COALESCE(SUM(t.paslon3), 0) AS paslon3,
			COALESCE(SUM(t.paslon4), 0) AS paslon4,
			COALESCE(SUM(t.suara_sah), 0) AS suara_sah,
			COALESCE(SUM(t.suara_tidak_sah), 0) AS suara_tidak_sah,
			COALESCE(k.total_voters, 0) AS total_voters,
			COALESCE(k.total_tps, 0) AS total_tps,
			COUNT(CASE WHEN t.suara_sah > 0 THEN 1 END) AS sudah,
			COUNT(CASE WHEN t.suara_sah = 0 THEN 1 END) AS belum,
			ROUND(
				(COALESCE(SUM(t.suara_sah), 0) + COALESCE(SUM(t.suara_tidak_sah), 0)) * 100.0 / COALESCE(NULLIF(k.total_voters, 0), 1)
			) AS pp,
			k.code
		FROM
			kelurahan k
		LEFT JOIN
			tps t ON k.kelurahan_id = t.kelurahan_id
		LEFT JOIN
			kecamatan c ON k.kecamatan_id = c.kecamatan_id
		WHERE
			c.code = $1
		GROUP BY
			k.kelurahan_id, k.kelurahan_name, c.kecamatan_name, k.total_voters, k.total_tps, k.code
		ORDER BY
			k.kelurahan_name;
	`

	err = r.db.SelectContext(ctx, &kelurahans, query, codeKecamatan)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}
	return

}

func (r repository) deleteDataKecamatanKelurahanTPS(ctx context.Context) (err error) {
	query := `
	DELETE FROM tps
	`

	if _, err = r.db.ExecContext(ctx, query); err != nil {
		return
	}

	query = `
	DELETE FROM kelurahan
	`

	if _, err = r.db.ExecContext(ctx, query); err != nil {
		return
	}

	query = `
	DELETE FROM kecamatan
	`

	if _, err = r.db.ExecContext(ctx, query); err != nil {
		return
	}

	query = `
	DELETE FROM auth WHERE role NOT IN ('admin', 'user')
	`

	if _, err = r.db.ExecContext(ctx, query); err != nil {
		return
	}

	return nil
}

func (r repository) createKecamatan(ctx context.Context, model Output) (err error) {
	query := `
		INSERT INTO kecamatan (
			kecamatan_id, kecamatan_name, total_voters, total_tps, code
		) VALUES (
			:kecamatan_id, :kecamatan_name, :total_voters, :total_tps, :code
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return
}

func (r repository) createKelurahan(ctx context.Context, model KelurahanOutput) (err error) {
	query := `
		INSERT INTO kelurahan (
			kelurahan_id, kecamatan_id, kelurahan_name, total_voters, total_tps, code
		) VALUES (
			:kelurahan_id, :kecamatan_id, :kelurahan_name, :total_voters_kelurahan, :total_tps_kelurahan, :code_kelurahan
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return nil
}

// Fungsi untuk menyisipkan data ke dalam tabel tps
func (r repository) createTps(ctx context.Context, model TpsOutput) error {
	query := `
		INSERT INTO tps (
			tps_id, kecamatan_id, kelurahan_id, tps_name, user_id, total_voters,
			paslon1, paslon2, paslon3, paslon4, suara_sah, suara_tidak_sah, photo, code
		) VALUES (
			:tps_id, :kecamatan_id, :kelurahan_id, :tps_name, :user_id, :total_voters,
			:paslon1, :paslon2, :paslon3, :paslon4, :suara_sah, :suara_tidak_sah, :photo, :code
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func (r repository) GetListKecamatanCode(ctx context.Context) (kecamatans []Kecamatan, err error) {
	query := `
	    SELECT
			kecamatan_name,
			code
		FROM
			kecamatan
		ORDER BY
			kecamatan_name ASC;
	`

	err = r.db.SelectContext(ctx, &kecamatans, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}
	return

}
