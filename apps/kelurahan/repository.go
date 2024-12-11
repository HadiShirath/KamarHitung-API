package kelurahan

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

func (r repository) GetListKelurahanCode(ctx context.Context, codeKecamatan string) (kelurahans []Kelurahan, err error) {
	query := `
	   SELECT
			k.kelurahan_name,
			k.code 
		FROM
			kelurahan k
		JOIN
			kecamatan c ON k.kecamatan_id = c.kecamatan_id
		WHERE
			c.code = $1
		ORDER BY
			k.kelurahan_name ASC
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

func (r repository) GetKelurahanData(ctx context.Context, codeKelurahan string) (kelurahan Kelurahan, err error) {
	query := `
		SELECT
    c.kecamatan_name,
    k.kelurahan_name,
    COALESCE(SUM(t.paslon1), 0) AS paslon1,
    COALESCE(SUM(t.paslon2), 0) AS paslon2,
    COALESCE(SUM(t.paslon3), 0) AS paslon3,
    COALESCE(SUM(t.paslon4), 0) AS paslon4,
	 COALESCE(SUM(t.suara_tidak_sah), 0) AS suara_tidak_sah
		FROM
			kelurahan k
		JOIN
			tps t ON k.kelurahan_id = t.kelurahan_id
		JOIN
			kecamatan c ON k.kecamatan_id = c.kecamatan_id
		WHERE
			k.code = $1
		GROUP BY
			c.kecamatan_name, k.kelurahan_name;
	`
	err = r.db.GetContext(ctx, &kelurahan, query, codeKelurahan)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) GetListTPSFromKelurahan(ctx context.Context, codeKelurahan string) (tpss []TPS, err error) {
	query := `
	SELECT
    k.kecamatan_name,
    l.kelurahan_name,
    t.tps_name,
    t.paslon1,
    t.paslon2,
    t.paslon3,
    t.paslon4,
    t.suara_sah,
    t.suara_tidak_sah,
    t.total_voters,
    CASE WHEN t.suara_sah > 0 THEN 1 ELSE 0 END AS sudah,
    CASE WHEN t.suara_sah = 0 THEN 1 ELSE 0 END AS belum,
    ROUND(
        (t.suara_sah + t.suara_tidak_sah) * 100.0 / NULLIF(t.total_voters, 0)
    ) AS pp,
    t.code
	FROM 
		tps t
	JOIN 
		kecamatan k ON t.kecamatan_id = k.kecamatan_id
	JOIN 
		kelurahan l ON t.kelurahan_id = l.kelurahan_id
	WHERE 
		l.code = $1
	ORDER BY 
		t.tps_name;
	`

	err = r.db.SelectContext(ctx, &tpss, query, codeKelurahan)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}
	return

}
