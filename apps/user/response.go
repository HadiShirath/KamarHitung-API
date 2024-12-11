package user

type UserListResponse struct {
	PublicID        string `json:"public_id"`
	Username        string `json:"username"`
	Fullname        string `json:"fullname"`
	PasswordDecoded string `json:"password_decoded"`
	Role            string `json:"role"`
}

func NewUserListResponseFromEntity(users []User) []UserListResponse {
	var userList = []UserListResponse{}

	for _, user := range users {
		userList = append(userList, user.ToUserListResponse())
	}

	return userList
}

type UserSaksiListResponse struct {
	Username      string `json:"username"`
	Fullname      string `json:"fullname"`
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	TpsName       string `json:"tps_name"`
}

func NewUserSaksiListResponseFromEntity(users []User) []UserSaksiListResponse {
	var userList = []UserSaksiListResponse{}

	for _, user := range users {
		userList = append(userList, user.ToUserSaksiListResponse())
	}

	return userList
}

type ExportDataCSVResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	TpsName       string `json:"tps_name"`
	FullName      string `json:"fullname"`
	Username      string `json:"username"`
	Password      string `json:"password_decoded"`
	CodeUnique    string `json:"code_unique"`
}

func NewExportDataCSVResponseFromEntity(users []User) []ExportDataCSVResponse {
	var exportDataCSVList = []ExportDataCSVResponse{}

	for _, user := range users {
		exportDataCSVList = append(exportDataCSVList, user.ToExportDataCSVResponse())
	}

	return exportDataCSVList
}
