package kecamatan

type GetVoterKecamatanResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	Paslon1       int32  `json:"paslon1"`
	Paslon2       int32  `json:"paslon2"`
	Paslon3       int32  `json:"paslon3"`
	Paslon4       int32  `json:"paslon4"`
	SuaraTidakSah int32  `json:"suara_tidak_sah"`
}

type AllVoterResponse struct {
	TotalSuara int32 `json:"total_suara"`
	Persentase int32 `json:"persentase"`
}

type KecamatanListResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	Paslon1       int32  `json:"paslon1"`
	Paslon2       int32  `json:"paslon2"`
	Paslon3       int32  `json:"paslon3"`
	Paslon4       int32  `json:"paslon4"`
	SuaraSah      int32  `json:"suara_sah"`
	SuaraTidakSah int32  `json:"suara_tidak_sah"`
	TotalVoters   int32  `json:"total_voters"`
	TotalTPS      int32  `json:"total_tps"`
	Sudah         int32  `json:"sudah"`
	Belum         int32  `json:"belum"`
	PP            int32  `json:"pp"`
	Code          string `json:"code"`
}

type KecamatanCodeListResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	Code          string `json:"code"`
}

func NewKecamatanCodeResponseFromEntity(kecamatans []Kecamatan) []KecamatanCodeListResponse {
	var kecamatanCodeList = []KecamatanCodeListResponse{}

	for _, kecamatan := range kecamatans {
		kecamatanCodeList = append(kecamatanCodeList, kecamatan.ToKecamatanCodeListResponse())
	}

	return kecamatanCodeList
}

func NewKecamatanListResponseFromEntity(Kecamatans []Kecamatan) []KecamatanListResponse {
	var KecamatanList = []KecamatanListResponse{}

	for _, Kecamatan := range Kecamatans {
		KecamatanList = append(KecamatanList, Kecamatan.ToKecamatanListResponse())
	}

	return KecamatanList
}

type KelurahanListResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	Paslon1       int32  `json:"paslon1"`
	Paslon2       int32  `json:"paslon2"`
	Paslon3       int32  `json:"paslon3"`
	Paslon4       int32  `json:"paslon4"`
	SuaraSah      int32  `json:"suara_sah"`
	SuaraTidakSah int32  `json:"suara_tidak_sah"`
	TotalVoters   int32  `json:"total_voters"`
	TotalTPS      int32  `json:"total_tps"`
	Sudah         int32  `json:"sudah"`
	Belum         int32  `json:"belum"`
	PP            int32  `json:"pp"`
	Code          string `json:"code"`
}

func NewKelurahanListResponseFromEntity(kelurahans []Kelurahan) []KelurahanListResponse {
	var kelurahanList = []KelurahanListResponse{}

	for _, kelurahan := range kelurahans {
		kelurahanList = append(kelurahanList, kelurahan.ToKelurahanListResponse())
	}

	return kelurahanList
}
