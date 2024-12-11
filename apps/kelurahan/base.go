package kelurahan

import (
	infrafiber "kamar-hitung/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	kelurahanRoute := router.Group("kelurahan")
	{
		kelurahanRoute.Get("/:code", infrafiber.CheckAuth(), handler.GetListKelurahanCode)
		kelurahanRoute.Get("/:code/detail", infrafiber.CheckAuth(), handler.GetListTPSFromKelurahan)
		kelurahanRoute.Get("/voter/:code", infrafiber.CheckAuth(), handler.GetKelurahanData)
	}
}
