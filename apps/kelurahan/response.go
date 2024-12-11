package kelurahan

type KelurahanDetailResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	Paslon1       int32  `json:"paslon1"`
	Paslon2       int32  `json:"paslon2"`
	Paslon3       int32  `json:"paslon3"`
	Paslon4       int32  `json:"paslon4"`
	SuaraTidakSah int32  `json:"suara_tidak_sah"`
}

type TPSListResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	TPSName       string `json:"tps_name"`
	Paslon1       int32  `json:"paslon1"`
	Paslon2       int32  `json:"paslon2"`
	Paslon3       int32  `json:"paslon3"`
	Paslon4       int32  `json:"paslon4"`
	SuaraSah      int32  `json:"suara_sah"`
	SuaraTidakSah int32  `json:"suara_tidak_sah"`
	TotalVoters   int32  `json:"total_voters"`
	Sudah         int32  `json:"sudah"`
	Belum         int32  `json:"belum"`
	PP            int32  `json:"pp"`
	Code          string `json:"code"`
}

func NewTPSListResponseFromEntity(tpss []TPS) []TPSListResponse {
	var tpsList = []TPSListResponse{}

	for _, kelurahan := range tpss {
		tpsList = append(tpsList, kelurahan.ToTPSListResponse())
	}

	return tpsList
}

type KelurahanCodeListResponse struct {
	KelurahanName string `json:"kelurahan_name"`
	Code          string `json:"code"`
}

func NewKelurahanCodeResponseFromEntity(kelurahans []Kelurahan) []KelurahanCodeListResponse {
	var kelurahanCodeList = []KelurahanCodeListResponse{}

	for _, kelurahan := range kelurahans {
		kelurahanCodeList = append(kelurahanCodeList, kelurahan.ToKelurahanCodeListResponse())
	}

	return kelurahanCodeList
}
