package tps

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

	tpsRoute := router.Group("tps")
	{
		tpsRoute.Get("/", infrafiber.CheckAuth(), handler.TPSAdressDetail)
		tpsRoute.Get("/all", infrafiber.CheckAuth(), handler.GetListTPS)
		tpsRoute.Get("/saksi", infrafiber.CheckAuth(), handler.GetTPSSaksiPagination)
		tpsRoute.Get("/:code", infrafiber.CheckAuth(), handler.GetListTPSCode)
		tpsRoute.Post("/photo", infrafiber.CheckAuth(), handler.CreatePhoto)
		tpsRoute.Post("/upload", infrafiber.CheckAuth(), handler.UploadDataTPS)
		tpsRoute.Get("/voter/all", infrafiber.CheckAuth(), handler.GetAllVoterTPS)
		tpsRoute.Get("/voter/:code", infrafiber.CheckAuth(), handler.GetVoterTPS)
		tpsRoute.Put("/voter/user/:id", infrafiber.CheckAuth(), handler.UpdateVoteTPSByUser)
		tpsRoute.Put("/voter/:id", infrafiber.CheckAuth(), handler.UpdateVoteTPS)
	}
}
