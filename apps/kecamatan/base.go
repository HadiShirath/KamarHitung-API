package kecamatan

import (
	"kamar-hitung/apps/auth"
	infrafiber "kamar-hitung/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	repoAuth := auth.NewRepository(db)
	svc := NewService(repo, repoAuth)
	handler := NewHandler(svc)

	kecamatanRoute := router.Group("kecamatan")
	{
		kecamatanRoute.Get("/", infrafiber.CheckAuth(), handler.GetListKecamatanCode)
		kecamatanRoute.Get("/all", infrafiber.CheckAuth(), handler.ListKecamatan)
		kecamatanRoute.Get("/voters", infrafiber.CheckAuth(), handler.GetAllVoter)
		kecamatanRoute.Get("/:code", infrafiber.CheckAuth(), handler.GetListKelurahanFromKecamatan)
		kecamatanRoute.Get("/voter/:code", infrafiber.CheckAuth(), handler.GetVoterKecamatan)
		kecamatanRoute.Post("/file/csv", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.importCSVHandler)
	}
}
