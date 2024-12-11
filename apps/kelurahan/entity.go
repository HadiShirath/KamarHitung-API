package kelurahan

type Kelurahan struct {
	KecamatanName string `db:"kecamatan_name"`
	KelurahanName string `db:"kelurahan_name"`
	Paslon1       int32  `db:"paslon1"`
	Paslon2       int32  `db:"paslon2"`
	Paslon3       int32  `db:"paslon3"`
	Paslon4       int32  `db:"paslon4"`
	SuaraSah      int32  `db:"suara_sah"`
	SuaraTidakSah int32  `db:"suara_tidak_sah"`
	TotalSuara    int32  `db:"total_suara"`
	Persentase    int32  `db:"persentase"`
	TotalVoters   int32  `db:"total_voters"`
	TotalTPS      int32  `db:"total_tps"`
	Sudah         int32  `db:"sudah"`
	Belum         int32  `db:"belum"`
	PP            int32  `db:"pp"`
	Code          string `db:"code"`
}

type TPS struct {
	KecamatanName string `db:"kecamatan_name"`
	KelurahanName string `db:"kelurahan_name"`
	TPSName       string `db:"tps_name"`
	Paslon1       int32  `db:"paslon1"`
	Paslon2       int32  `db:"paslon2"`
	Paslon3       int32  `db:"paslon3"`
	Paslon4       int32  `db:"paslon4"`
	SuaraSah      int32  `db:"suara_sah"`
	SuaraTidakSah int32  `db:"suara_tidak_sah"`
	TotalSuara    int32  `db:"total_suara"`
	Persentase    int32  `db:"persentase"`
	TotalVoters   int32  `db:"total_voters"`
	TotalTPS      int32  `db:"total_tps"`
	Sudah         int32  `db:"sudah"`
	Belum         int32  `db:"belum"`
	PP            int32  `db:"pp"`
	Code          string `db:"code"`
}

func (k Kelurahan) ToKelurahanCodeListResponse() KelurahanCodeListResponse {
	return KelurahanCodeListResponse{
		KelurahanName: k.KelurahanName,
		Code:          k.Code,
	}
}

func (k Kelurahan) ToGetVoterKelurahanResponse() KelurahanDetailResponse {
	return KelurahanDetailResponse{
		KecamatanName: k.KecamatanName,
		KelurahanName: k.KelurahanName,
		Paslon1:       k.Paslon1,
		Paslon2:       k.Paslon2,
		Paslon3:       k.Paslon3,
		Paslon4:       k.Paslon4,
		SuaraTidakSah: k.SuaraTidakSah,
	}
}

func (t TPS) ToTPSListResponse() TPSListResponse {
	return TPSListResponse{
		KecamatanName: t.KecamatanName,
		KelurahanName: t.KelurahanName,
		TPSName:       t.TPSName,
		Paslon1:       t.Paslon1,
		Paslon2:       t.Paslon2,
		Paslon3:       t.Paslon3,
		Paslon4:       t.Paslon4,
		SuaraSah:      t.SuaraSah,
		SuaraTidakSah: t.SuaraTidakSah,
		TotalVoters:   t.TotalVoters,
		Sudah:         t.Sudah,
		Belum:         t.Belum,
		PP:            t.PP,
		Code:          t.Code,
	}
}
