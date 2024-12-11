package infrafiber

import (
	"fmt"
	"kamar-hitung/infra/response"
	"kamar-hitung/internal/config"
	"kamar-hitung/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		if authorization == "" {
			return NewResponse(
				WithMessage("unauthorized"),
				WithError(response.ErrorUnAuthorized),
			).Send(c)
		}

		bearer := strings.Split(authorization, "Bearer ")

		if len(bearer) != 2 {
			return NewResponse(
				WithMessage("token invalid"),
				WithError(response.ErrorUnAuthorized),
			).Send(c)
		}

		token := bearer[1]
		publicId, role, err := utility.ValidateToken(token, config.Cfg.App.Encryption.JWTSecret)
		if err != nil {
			return NewResponse(
				WithMessage(err.Error()),
				WithError(response.ErrorUnAuthorized),
			).Send(c)
		}

		c.Locals("ROLE", role)
		c.Locals("PUBLIC_ID", publicId)

		return c.Next()
	}
}

func CheckRoles(authorizedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := fmt.Sprintf("%+v", c.Locals("ROLE"))

		isExist := false

		for _, authorizedRole := range authorizedRoles {
			if role == authorizedRole {
				isExist = true
				break
			}
		}

		if !isExist {
			requiredRoles := strings.Join(authorizedRoles, " & ")

			return NewResponse(
				WithMessage(fmt.Sprintf("access %s only", requiredRoles)),
				WithError(response.ErrorForbiddenAccess),
			).Send(c)
		}

		return c.Next()
	}
}
