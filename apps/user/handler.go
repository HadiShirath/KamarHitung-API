package user

import (
	"context"
	infrafiber "kamar-hitung/infra/fiber"
	"kamar-hitung/infra/response"
	"net/http"

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

func (h handler) GetUserList(ctx *fiber.Ctx) error {

	users, err := h.svc.GetUserList(context.Background())

	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	userListResponse := NewUserListResponseFromEntity(users)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list user success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(userListResponse),
	).Send(ctx)
}

func (h handler) GetUserSaksiList(ctx *fiber.Ctx) error {

	users, err := h.svc.GetUserSaksiList(context.Background())

	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	userListResponse := NewUserSaksiListResponseFromEntity(users)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list user success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(userListResponse),
	).Send(ctx)
}

func (h handler) EditAuth(ctx *fiber.Ctx) error {
	var req = EditUserRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage("register failed"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	userId := ctx.Params("id", "")
	if userId == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	err := h.svc.EditTPSSaksi(context.Background(), req, userId)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("update user success"),
	).Send(ctx)
}

func (h handler) GetDataForExportCSV(ctx *fiber.Ctx) error {

	data, err := h.svc.GetDataForExportCSV(context.Background())
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	getExportDataList := NewExportDataCSVResponseFromEntity(data)

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get data saksi succesfully"),
		infrafiber.WithPayload(getExportDataList),
	).Send(ctx)

}

func (h handler) DeleteUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id", "")
	if userId == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	err := h.svc.DeleteUserById(context.Background(), userId)

	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("delete user success"),
	).Send(ctx)
}
