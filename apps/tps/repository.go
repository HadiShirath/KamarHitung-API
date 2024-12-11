package tps

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

func (r repository) GetListTPSCode(ctx context.Context, codeKelurahan string) (tpss []TPS, err error) {
	query := `
	   SELECT
			t.tps_name,
			t.code 
		FROM
			tps t
		JOIN
			kelurahan k ON t.kelurahan_id = k.kelurahan_id
		WHERE
			k.code = $1
		ORDER BY
			t.tps_name ASC
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

func (r repository) CreatePhoto(ctx context.Context, model TPS, userId string) (err error) {
	query := `
		UPDATE tps SET photo=$1 WHERE user_id=$2
	`

	_, err = r.db.ExecContext(ctx, query, model.Photo, userId)

	if err != nil {
		return
	}

	return
}

func (r repository) UploadDataTPS(ctx context.Context, model TPS, userId string) (err error) {

	if model.Photo != "" {
		query := `
		UPDATE tps SET paslon1=$1, paslon2=$2, paslon3=$3, paslon4=$4, suara_sah=$5, suara_tidak_sah=$6, photo=$7 WHERE user_id=$8
		`

		_, err = r.db.ExecContext(ctx, query, model.Paslon1, model.Paslon2, model.Paslon3, model.Paslon4, model.SuaraSah, model.SuaraTidakSah, model.Photo, userId)
	} else {
		query := `
		UPDATE tps SET paslon1=$1, paslon2=$2, paslon3=$3, paslon4=$4, suara_sah=$5, suara_tidak_sah=$6 WHERE user_id=$7
		`

		_, err = r.db.ExecContext(ctx, query, model.Paslon1, model.Paslon2, model.Paslon3, model.Paslon4, model.SuaraSah, model.SuaraTidakSah, userId)
	}

	if err != nil {
		return
	}

	return
}

func (r repository) GetAddressTPSByUserId(ctx context.Context, userId string) (tps TPS, err error) {
	query := `
		SELECT
		  tps.user_id, kecamatan.kecamatan_name, kelurahan.kelurahan_name , tps.tps_name, tps.paslon1, tps.paslon2, tps.paslon3, tps.paslon4, tps.suara_sah, tps.suara_tidak_sah, tps.photo, auth.fullname
		FROM tps
		INNER JOIN kecamatan ON tps.kecamatan_id = kecamatan.kecamatan_id
		INNER JOIN kelurahan ON tps.kelurahan_id = kelurahan.kelurahan_id
		INNER JOIN auth ON tps.user_id = auth.public_id
    	WHERE tps.user_id=$1
	`

	err = r.db.GetContext(ctx, &tps, query, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) GetAllVoterTPS(ctx context.Context) (tps TPS, err error) {
	query := `
		  SELECT
			COALESCE(SUM(t.paslon1), 0) as paslon1,
			COALESCE(SUM(t.paslon2), 0) AS paslon2,
			COALESCE(SUM(t.paslon3), 0) AS paslon3,
			COALESCE(SUM(t.paslon4), 0) AS paslon4,
			COALESCE(SUM(t.suara_tidak_sah), 0) AS suara_tidak_sah
		FROM
    		tps t
	`
	err = r.db.GetContext(ctx, &tps, query)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) GetVoterTPS(ctx context.Context, codeTPS string) (tps TPS, err error) {
	query := `
		SELECT
			kecamatan.kecamatan_name, kelurahan.kelurahan_name , tps.tps_name, tps.paslon1, tps.paslon2, tps.paslon3, tps.paslon4, tps.suara_sah, tps.suara_tidak_sah, tps.photo
		FROM tps
		INNER JOIN kecamatan ON tps.kecamatan_id = kecamatan.kecamatan_id
		INNER JOIN kelurahan ON tps.kelurahan_id = kelurahan.kelurahan_id
    	WHERE tps.code=$1
	`
	err = r.db.GetContext(ctx, &tps, query, codeTPS)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) GetAllTPSSaksiWithPaginationCursor(ctx context.Context, model TPSPagination) (tpss []TPS, err error) {
	query := `
	SELECT
    k.kecamatan_name,
    l.kelurahan_name,
    t.tps_name,
	a.public_id as user_id,
    a.fullname as name_koordinator,
    a.username as hp_koordinator,
    t.code
		FROM tps t
	JOIN kelurahan l ON t.kelurahan_id = l.kelurahan_id
	JOIN kecamatan k ON l.kecamatan_id = k.kecamatan_id
	JOIN auth a ON t.user_id = a.public_id
		ORDER BY k.kecamatan_name ASC, l.kelurahan_name ASC, t.tps_name ASC
	`

	// err = r.db.SelectContext(ctx, &tpss, query, model.Offset, model.Limit)
	err = r.db.SelectContext(ctx, &tpss, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return
	}

	return
}

func (r repository) EditTPSSaksi(ctx context.Context, model TPS, userId string) (err error) {
	query := `
		UPDATE auth SET fullname=$1, username=$2 WHERE user_id=$2
	`

	_, err = r.db.ExecContext(ctx, query, model.Fullname, model.Username, userId)

	if err != nil {
		return
	}

	return
}

func (r repository) GetListTPS(ctx context.Context) (tpss []TPS, err error) {
	query := `
	SELECT
		k.kecamatan_name,
		l.kelurahan_name,
		t.tps_id,
		t.tps_name,
		t.paslon1,
		t.paslon2,
		t.paslon3,
		t.paslon4,
		t.suara_sah,
		t.suara_tidak_sah,
		t.total_voters,
		ROUND(
			(t.suara_sah + t.suara_tidak_sah) * 100.0 / NULLIF(t.total_voters, 0)
		) AS pp,
		t.code,
		t.photo
	FROM 
		tps t
	JOIN 
		kecamatan k ON t.kecamatan_id = k.kecamatan_id
	JOIN 
		kelurahan l ON t.kelurahan_id = l.kelurahan_id
	ORDER BY 
		k.kecamatan_name ASC, l.kelurahan_name ASC, t.tps_name ASC
	`

	err = r.db.SelectContext(ctx, &tpss, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}
	return

}

func (r repository) EditVoteTPS(ctx context.Context, model TPS, tpsId string) (err error) {
	query := `
		UPDATE tps SET paslon1=$1, paslon2=$2, paslon3=$3, paslon4=$4, suara_sah=$5, suara_tidak_sah=$6 WHERE tps_id=$7
	`

	_, err = r.db.ExecContext(ctx, query, model.Paslon1, model.Paslon2, model.Paslon3, model.Paslon4, model.SuaraSah, model.SuaraTidakSah, tpsId)

	if err != nil {
		return
	}

	return
}

func (r repository) UpdateVoteTPSByUserId(ctx context.Context, model TPS, userId string) (updatedModel TPS, err error) {
	query := `
		UPDATE tps
		SET paslon1=$1, paslon2=$2, paslon3=$3, paslon4=$4, suara_sah=$5, suara_tidak_sah=$6
		WHERE user_id=$7
	`

	_, err = r.db.ExecContext(ctx, query, model.Paslon1, model.Paslon2, model.Paslon3, model.Paslon4, model.SuaraSah, model.SuaraTidakSah, userId)
	if err != nil {
		return updatedModel, err
	}

	// Fetch updated data
	selectQuery := `
		SELECT
		  kecamatan.kecamatan_name, kelurahan.kelurahan_name, tps.tps_name, tps.photo
		FROM tps
		INNER JOIN kecamatan ON tps.kecamatan_id = kecamatan.kecamatan_id
		INNER JOIN kelurahan ON tps.kelurahan_id = kelurahan.kelurahan_id
		INNER JOIN auth ON tps.user_id = auth.public_id
    	WHERE tps.user_id=$1
	`

	err = r.db.GetContext(ctx, &updatedModel, selectQuery, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}
