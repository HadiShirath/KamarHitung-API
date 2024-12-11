package tps

type TpsDetailResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	TpsName       string `json:"tps_name"`
	Paslon1       int32  `json:"paslon1"`
	Paslon2       int32  `json:"paslon2"`
	Paslon3       int32  `json:"paslon3"`
	Paslon4       int32  `json:"paslon4"`
	SuaraSah      int32  `json:"suara_sah"`
	SuaraTidakSah int32  `json:"suara_tidak_sah"`
	Photo         string `json:"photo"`
	Fullname      string `json:"fullname"`
}
type TpsDetailResponseFromUpdateData struct {
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	TpsName       string `json:"tps_name"`
}

type GetVoterTPSResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	TpsName       string `json:"tps_name"`
	Paslon1       int32  `json:"paslon1"`
	Paslon2       int32  `json:"paslon2"`
	Paslon3       int32  `json:"paslon3"`
	Paslon4       int32  `json:"paslon4"`
	SuaraSah      int32  `json:"suara_sah"`
	SuaraTidakSah int32  `json:"suara_tidak_sah"`
	Photo         string `json:"photo"`
}

type GetAllVoterTPSResponse struct {
	Paslon1       int32 `json:"paslon1"`
	Paslon2       int32 `json:"paslon2"`
	Paslon3       int32 `json:"paslon3"`
	Paslon4       int32 `json:"paslon4"`
	SuaraTidakSah int32 `json:"suara_tidak_sah"`
}

type TPSListSaksiResponse struct {
	KecamatanName   string `json:"kecamatan_name"`
	KelurahanName   string `json:"kelurahan_name"`
	TpsName         string `json:"tps_name"`
	UserId          string `json:"user_id"`
	NameKoordinator string `json:"name_koordinator"`
	HpKoordinator   string `json:"hp_koordinator"`
	Code            string `json:"code"`
}

type TPSCodeListResponse struct {
	TpsName string `json:"tps_name"`
	Code    string `json:"code"`
}

func NewTPSCodeResponseFromEntity(tpss []TPS) []TPSCodeListResponse {
	var tpsCodeList = []TPSCodeListResponse{}

	for _, kelurahan := range tpss {
		tpsCodeList = append(tpsCodeList, kelurahan.ToTPSCodeListResponse())
	}

	return tpsCodeList
}

func NewTPSListSaksiResponseFromEntity(tpss []TPS) []TPSListSaksiResponse {
	var TPSListSaksi = []TPSListSaksiResponse{}

	for _, saksi := range tpss {
		TPSListSaksi = append(TPSListSaksi, saksi.ToTPSListSaksiResponse())
	}

	return TPSListSaksi
}

type TPSListResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	TpsId         string `json:"tps_id"`
	TPSName       string `json:"tps_name"`
	Paslon1       int32  `json:"paslon1"`
	Paslon2       int32  `json:"paslon2"`
	Paslon3       int32  `json:"paslon3"`
	Paslon4       int32  `json:"paslon4"`
	SuaraSah      int32  `json:"suara_sah"`
	SuaraTidakSah int32  `json:"suara_tidak_sah"`
	TotalVoters   int32  `json:"total_voters"`
	Photo         string `json:"photo"`
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
