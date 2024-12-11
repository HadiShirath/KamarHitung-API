package tps

import (
	"kamar-hitung/infra/response"
	"kamar-hitung/internal/config"
	"kamar-hitung/utility"
	"log"
	"strconv"
)

type TPS struct {
	UserId          string `db:"user_id"`
	KecamatanName   string `db:"kecamatan_name"`
	KelurahanName   string `db:"kelurahan_name"`
	TpsId           string `db:"tps_id"`
	TpsName         string `db:"tps_name"`
	Photo           string `db:"photo"`
	Fullname        string `db:"fullname"`
	Username        string `db:"username"`
	NameKoordinator string `db:"name_koordinator"`
	HpKoordinator   string `db:"hp_koordinator"`
	Paslon1         int32  `db:"paslon1"`
	Paslon2         int32  `db:"paslon2"`
	Paslon3         int32  `db:"paslon3"`
	Paslon4         int32  `db:"paslon4"`
	SuaraSah        int32  `db:"suara_sah"`
	SuaraTidakSah   int32  `db:"suara_tidak_sah"`
	TotalVoters     int32  `db:"total_voters"`
	PP              int32  `db:"pp"`
	Code            string `db:"code"`
}

type TPSPagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func NewFromCreatePhotoRequest(req CreatePhotoRequestPayload) TPS {
	return TPS{
		Photo: req.Photo,
	}
}

func NewFromUploadDataRequest(req UploadDataRequestPayload) TPS {
	return TPS{
		Paslon1:       parseToInt32(req.Paslon1),
		Paslon2:       parseToInt32(req.Paslon2),
		Paslon3:       parseToInt32(req.Paslon3),
		Paslon4:       parseToInt32(req.Paslon4),
		SuaraSah:      parseToInt32(req.SuaraSah),
		SuaraTidakSah: parseToInt32(req.SuaraTidakSah),
		Photo:         req.Photo,
	}
}

func NewFromEditTPSSaksiRequest(req EditTPSSaksiRequestPayload) TPS {
	return TPS{
		Fullname: req.Fullname,
		Username: req.Username,
		UserId:   req.UserId,
	}
}

func NewFromEditVoteTPSRequest(req EditVoteTPSRequestPayload) TPS {
	return TPS{
		Paslon1:       req.Paslon1,
		Paslon2:       req.Paslon2,
		Paslon3:       req.Paslon3,
		Paslon4:       req.Paslon4,
		SuaraSah:      req.SuaraSah,
		SuaraTidakSah: req.SuaraTidakSah,
	}
}

func NewFromEditVoteBySaksiTPSRequest(req EditVoteTPSBySaksiRequestPayload) TPS {

	return TPS{
		Paslon1:       req.Paslon1,
		Paslon2:       req.Paslon2,
		Paslon3:       req.Paslon3,
		Paslon4:       req.Paslon4,
		SuaraSah:      req.Paslon1 + req.Paslon2 + req.Paslon3 + req.Paslon4,
		SuaraTidakSah: req.SuaraTidakSah,
	}
}

func NewTPSSaksiPaginationFromProductRequest(req ListTPSSaksiRequestPayload) TPSPagination {
	req = req.GenerateDefaultValue()
	return TPSPagination{
		Offset: req.Offset,
		Limit:  req.Limit,
	}
}

func (t TPS) ToTPSCodeListResponse() TPSCodeListResponse {
	return TPSCodeListResponse{
		TpsName: t.TpsName,
		Code:    t.Code,
	}
}

func (t TPS) ToTPSListSaksiResponse() TPSListSaksiResponse {
	return TPSListSaksiResponse{
		KecamatanName:   t.KecamatanName,
		KelurahanName:   t.KelurahanName,
		TpsName:         t.TpsName,
		UserId:          t.UserId,
		NameKoordinator: t.NameKoordinator,
		HpKoordinator:   t.HpKoordinator,
		Code:            t.Code,
	}
}

func (t TPS) ToTPSListResponse() TPSListResponse {
	return TPSListResponse{
		KecamatanName: t.KecamatanName,
		KelurahanName: t.KelurahanName,
		TpsId:         t.TpsId,
		TPSName:       t.TpsName,
		Paslon1:       t.Paslon1,
		Paslon2:       t.Paslon2,
		Paslon3:       t.Paslon3,
		Paslon4:       t.Paslon4,
		SuaraSah:      t.SuaraSah,
		SuaraTidakSah: t.SuaraTidakSah,
		TotalVoters:   t.TotalVoters,
		Photo:         t.Photo,
		PP:            t.PP,
		Code:          t.Code,
	}
}

func (t TPS) ToTpsDetailResponse() TpsDetailResponse {
	return TpsDetailResponse{
		KecamatanName: t.KecamatanName,
		KelurahanName: t.KelurahanName,
		TpsName:       t.TpsName,
		Paslon1:       t.Paslon1,
		Paslon2:       t.Paslon2,
		Paslon3:       t.Paslon3,
		Paslon4:       t.Paslon4,
		SuaraSah:      t.SuaraSah,
		SuaraTidakSah: t.SuaraTidakSah,
		Photo:         t.Photo,
		Fullname:      t.Fullname,
	}
}

func (t TPS) ToTpsDetailFromUpdateDataResponse() TpsDetailResponseFromUpdateData {
	return TpsDetailResponseFromUpdateData{
		KecamatanName: t.KecamatanName,
		KelurahanName: t.KelurahanName,
		TpsName:       t.TpsName,
	}
}

func (t TPS) ToGetAllVoterTPSResponse() GetAllVoterTPSResponse {
	return GetAllVoterTPSResponse{
		Paslon1:       t.Paslon1,
		Paslon2:       t.Paslon2,
		Paslon3:       t.Paslon3,
		Paslon4:       t.Paslon4,
		SuaraTidakSah: t.SuaraTidakSah,
	}
}

func (t TPS) ToGetVoterTPSResponse() GetVoterTPSResponse {
	return GetVoterTPSResponse{
		KecamatanName: t.KecamatanName,
		KelurahanName: t.KelurahanName,
		TpsName:       t.TpsName,
		Paslon1:       t.Paslon1,
		Paslon2:       t.Paslon2,
		Paslon3:       t.Paslon3,
		Paslon4:       t.Paslon4,
		SuaraSah:      t.SuaraSah,
		SuaraTidakSah: t.SuaraTidakSah,
		Photo:         t.Photo,
	}
}

func (t TPS) ValidateUserId() (err error) {
	if len(t.UserId) != 36 {
		return response.ErrUserIdInvalid
	}

	return
}

func (t TPS) GenerateTokenData(secret string) (tokenString string, err error) {
	return utility.GenerateTokenData(t.UserId, t.KecamatanName, t.KelurahanName, t.TpsName, t.Photo, t.Fullname, secret)
}

func validateCodeUnique(codeUnique string) (err error) {
	if codeUnique != config.Cfg.App.Code {
		return response.ErrCodeInvalid
	}

	return
}

func parseToInt32(value string) int32 {
	parsedValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Println("error converting value")
	}
	return int32(parsedValue)
}

func (t TPS) ValidateSuaraSah() (err error) {
	if t.Paslon1+t.Paslon2+t.Paslon3+t.Paslon4 > 600+600*0.02 {
		return response.ErrTotalVoteInvalid
	}

	return
}
