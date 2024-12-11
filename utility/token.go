package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenData struct {
	ID        string `json:"id"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
	TPS       string `json:"tps"`
	Photo     string `json:"photo"`
	FullName  string `json:"fullname"`
}

func GenerateToken(id string, fullname string, role string, secret string) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"id":       id,
		"fullname": fullname,
		"role":     role,
		// "exp":      time.Now().Add(10 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string, secret string) (id string, role string, err error) {

	tokens, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return
	}

	claims, ok := tokens.Claims.(jwt.MapClaims)

	if ok && tokens.Valid {
		id = fmt.Sprintf("%+v", claims["id"])
		role = fmt.Sprintf("%+v", claims["role"])
		return
	}

	err = fmt.Errorf("unable to extract claims")
	return
}

func GenerateTokenData(id string, kecamatan string, kelurahan string, tps string, photo string, fullname string, secret string) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"id":        id,
		"kecamatan": kecamatan,
		"kelurahan": kelurahan,
		"tps":       tps,
		"photo":     photo,
		"fullname":  fullname,
		"exp":       time.Now().Add(10 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateTokenData(tokenString string, secret string) (data TokenData, err error) {

	tokens, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return
	}

	claims, ok := tokens.Claims.(jwt.MapClaims)

	if ok && tokens.Valid {
		token := TokenData{
			ID:        fmt.Sprintf("%+v", claims["id"]),
			Kecamatan: fmt.Sprintf("%+v", claims["kecamatan"]),
			Kelurahan: fmt.Sprintf("%+v", claims["kelurahan"]),
			TPS:       fmt.Sprintf("%+v", claims["tps"]),
			Photo:     fmt.Sprintf("%+v", claims["photo"]),
			FullName:  fmt.Sprintf("%+v", claims["fullname"]),
		}

		return token, nil
	}

	err = fmt.Errorf("unable to extract claims")
	return
}
