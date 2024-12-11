package kecamatan

import (
	"kamar-hitung/infra/response"
)

type Kecamatan struct {
	KecamatanName string `db:"kecamatan_name"`
	KelurahanName string `db:"kelurahan_name"`
	TpsName       string `db:"tps_name"`
	Fullname      string `db:"fullname"`
	Username      string `db:"username"`
	Password      string `db:"password_decoded"`
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

// Definisikan struct Data
type Data struct {
	KecamatanName string
	KelurahanName string
	TpsName       string
	TotalVoters   int
	Fullname      string
	Username      string
	Password      string
}

type Output struct {
	KecamatanID          string `db:"kecamatan_id"`
	KecamatanName        string `db:"kecamatan_name"`
	TotalVotersKecamatan int    `db:"total_voters"`
	TotalTpsKecamatan    int    `db:"total_tps"`
	CodeKecamatan        string `db:"code"`
	KelurahanID          string `db:"kelurahan_id"`
	KelurahanName        string `db:"kelurahan_name"`
	TotalVotersKelurahan int    `db:"total_voters_kelurahan"`
	TotalTpsKelurahan    int    `db:"total_tps_kelurahan"`
	CodeKelurahan        string `db:"code_kelurahan"`
	TpsID                string `db:"tps_id"`
	TpsName              string `db:"tps_name"`
	TotalVotersTPS       int    `db:"total_voters_tps"`
}

type KelurahanOutput struct {
	KelurahanID          string `db:"kelurahan_id"`
	KecamatanID          string `db:"kecamatan_id"`
	KelurahanName        string `db:"kelurahan_name"`
	TotalVotersKelurahan int    `db:"total_voters_kelurahan"`
	TotalTpsKelurahan    int    `db:"total_tps_kelurahan"`
	CodeKelurahan        string `db:"code_kelurahan"`
}

type TpsOutput struct {
	TpsID         string `db:"tps_id"`
	KecamatanID   string `db:"kecamatan_id"`
	KelurahanID   string `db:"kelurahan_id"`
	TpsName       string `db:"tps_name"`
	UserID        string `db:"user_id"`
	TotalVoters   int    `db:"total_voters"`
	Paslon1       int    `db:"paslon1"`
	Paslon2       int    `db:"paslon2"`
	Paslon3       int    `db:"paslon3"`
	Paslon4       int    `db:"paslon4"`
	SuaraSah      int    `db:"suara_sah"`
	SuaraTidakSah int    `db:"suara_tidak_sah"`
	Photo         string `db:"photo"`
	Code          string `db:"code"`
}

func (k Kecamatan) ToKecamatanListResponse() KecamatanListResponse {
	return KecamatanListResponse{
		KecamatanName: k.KecamatanName,
		Paslon1:       k.Paslon1,
		Paslon2:       k.Paslon2,
		Paslon3:       k.Paslon3,
		Paslon4:       k.Paslon4,
		SuaraSah:      k.SuaraSah,
		SuaraTidakSah: k.SuaraTidakSah,
		TotalVoters:   k.TotalVoters,
		TotalTPS:      k.TotalTPS,
		Sudah:         k.Sudah,
		Belum:         k.Belum,
		PP:            k.PP,
		Code:          k.Code,
	}
}

func (k Kecamatan) ToKecamatanCodeListResponse() KecamatanCodeListResponse {
	return KecamatanCodeListResponse{
		KecamatanName: k.KecamatanName,
		Code:          k.Code,
	}
}

func (k Kelurahan) ToKelurahanListResponse() KelurahanListResponse {
	return KelurahanListResponse{
		KecamatanName: k.KecamatanName,
		KelurahanName: k.KelurahanName,
		Paslon1:       k.Paslon1,
		Paslon2:       k.Paslon2,
		Paslon3:       k.Paslon3,
		Paslon4:       k.Paslon4,
		SuaraSah:      k.SuaraSah,
		SuaraTidakSah: k.SuaraTidakSah,
		TotalVoters:   k.TotalVoters,
		TotalTPS:      k.TotalTPS,
		Sudah:         k.Sudah,
		Belum:         k.Belum,
		PP:            k.PP,
		Code:          k.Code,
	}
}

func (k Kecamatan) ToGetVoterKecamatanResponse() GetVoterKecamatanResponse {
	return GetVoterKecamatanResponse{
		KecamatanName: k.KecamatanName,
		Paslon1:       k.Paslon1,
		Paslon2:       k.Paslon2,
		Paslon3:       k.Paslon3,
		Paslon4:       k.Paslon4,
		SuaraTidakSah: k.SuaraTidakSah,
	}
}

func (k Kecamatan) ToAllVoterResponse() AllVoterResponse {
	return AllVoterResponse{
		TotalSuara: k.TotalSuara,
		Persentase: k.Persentase,
	}
}

func validateFormatCSV(data []string) (err error) {
	expected := []string{"Kecamatan", "Kelurahan", "TPS", "DPT"}

	if len(data) != 6 {
		return response.ErrFormatCSVInvalid
	}

	for index, record := range data {
		if index < len(expected) && record != expected[index] {
			return response.ErrFormatCSVInvalid
		}
	}

	return
}
