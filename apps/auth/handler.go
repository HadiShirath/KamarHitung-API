package auth

import (
	infrafiber "kamar-hitung/infra/fiber"
	"kamar-hitung/infra/response"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func NewHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) register(ctx *fiber.Ctx) error {
	var req = RegisterRequestPayload{}
	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage("register failed"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	if _, err := h.svc.register(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("register failed"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("register success"),
	).Send(ctx)
}

func (h handler) login(ctx *fiber.Ctx) error {
	var req = LoginRequestPayload{}
	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage("login failed"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	token, role, err := h.svc.login(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	// Set cookie dengan atribut yang sesuai
	if role != "admin" {
		ctx.Cookie(&fiber.Cookie{
			Name:     "access_token",                   // Nama cookie
			Value:    token,                            // Nilai cookie
			Expires:  time.Now().Add(10 * time.Minute), // Waktu kedaluwarsa cookie
			SameSite: "Lax",                            // "Lax", "Strict", atau "None" sesuai kebutuhan
			Secure:   false,                            // Gunakan true jika aplikasi menggunakan HTTPS
			// Secure: true, // Gunakan true jika aplikasi menggunakan HTTPS
			// Domain: "kamarhitung.id",
		})
	} else {
		ctx.Cookie(&fiber.Cookie{
			Name:     "access_token",                   // Nama cookie
			Value:    token,                            // Nilai cookie
			SameSite: "Lax",                            // "Lax", "Strict", atau "None" sesuai kebutuhan
			Expires:  time.Now().Add(10 * time.Minute), // Waktu kedaluwarsa cookie
			Secure:   false,                            // Gunakan true jika aplikasi menggunakan HTTPS
			// Secure: true, // Gunakan true jika aplikasi menggunakan HTTPS
			// Domain: "kamarhitung.id",
		})
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("login success"),
		// infrafiber.WithPayload(map[string]interface{}{
		// 	"access_token": token,
		// 	"role":         role,
		// }),
	).Send(ctx)
}
