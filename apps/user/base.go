package user

import (
	"kamar-hitung/apps/auth"
	infrafiber "kamar-hitung/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	userRoute := router.Group("users")
	{
		userRoute.Get("/", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.GetUserList)
		userRoute.Get("/saksi", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.GetUserSaksiList)
		userRoute.Get("/saksi/csv", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.GetDataForExportCSV)
		userRoute.Put("/:id", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.EditAuth)
		userRoute.Delete("/:id", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.DeleteUserById)
	}
}
