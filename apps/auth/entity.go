package auth

import (
	"kamar-hitung/infra/response"
	"kamar-hitung/utility"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	ROLE_Admin Role = "admin"
	ROLE_User  Role = "user"
	ROLE_Saksi Role = "saksi"
)

type AuthEntity struct {
	PublicId        uuid.UUID `db:"public_id"`
	Username        string    `db:"username"`
	Fullname        string    `db:"fullname"`
	Password        string    `db:"password"`
	PasswordDecoded string    `db:"password_decoded"`
	Role            string    `db:"role"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

func NewFromRegisterRequest(req RegisterRequestPayload) AuthEntity {
	return AuthEntity{
		PublicId:        uuid.New(),
		Username:        req.Username,
		Fullname:        req.Fullname,
		Password:        req.Password,
		PasswordDecoded: req.Password,
		Role:            req.Role,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func NewFromLoginRequest(req LoginRequestPayload) AuthEntity {
	return AuthEntity{
		Username: req.Username,
		Password: req.Password,
	}
}

func (a AuthEntity) Validate() (err error) {
	if err = a.ValidateEmail(); err != nil {
		return
	}
	if err = a.ValidatePassword(); err != nil {
		return
	}

	return
}

func (a AuthEntity) ValidateEmail() (err error) {
	// if a.Username == "" {
	// 	return response.ErrEmailRequired
	// }

	// emails := strings.Split(a.Username, "@")
	// if len(emails) != 2 {
	// 	return response.ErrEmailInvalid
	// }

	return
}

func (a AuthEntity) ValidatePassword() (err error) {
	if a.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(a.Password) < 4 {
		return response.ErrPasswordInvalidLength
	}

	return
}

func (a AuthEntity) IsExists() bool {
	return a.PublicId.ID() != 0
}

func (a *AuthEntity) ValidateRole() (err error) {
	if a.Role == "" {
		a.Role = string(ROLE_Saksi)
	}

	return nil
}

func (a *AuthEntity) EncryptPassword(salt int) (err error) {
	encryptPass, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	a.Password = string(encryptPass)
	return nil
}

func (a AuthEntity) VerifyPasswordFromEncrypted(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
}

func (a AuthEntity) VerifyPasswordFromPlain(encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(a.Password))
}

func (a AuthEntity) GenerateToken(secret string) (tokenString string, err error) {
	return utility.GenerateToken(a.PublicId.String(), a.Fullname, string(a.Role), secret)
}
