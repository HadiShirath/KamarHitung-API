package message

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

	messageRoute := router.Group("messages")
	{
		messageRoute.Get("/inbox", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.GetInboxList)
		messageRoute.Get("/outbox", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.GetOutboxList)
		messageRoute.Post("/inbox", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.UploadInbox)
		messageRoute.Post("/outbox", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.CreateMessage)
		messageRoute.Post("/outboxs", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.CreateMessages)
		messageRoute.Put("/outbox/:id", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.UpdateStatusOutbox)
	}
}
