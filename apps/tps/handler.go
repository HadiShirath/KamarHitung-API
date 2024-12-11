package tps

import (
	"context"
	"fmt"
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

func (h handler) GetListTPSCode(ctx *fiber.Ctx) error {
	codeKelurahan := ctx.Params("code", "")
	if codeKelurahan == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	tpss, err := h.svc.GetListTPSCode(context.Background(), codeKelurahan)
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

	tpsCodeListResponse := NewTPSCodeResponseFromEntity(tpss)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list tps code success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(tpsCodeListResponse),
	).Send(ctx)
}

func (h handler) CreatePhoto(ctx *fiber.Ctx) error {

	user_id := fmt.Sprintf("%+v", ctx.Locals("PUBLIC_ID"))

	file, errFile := ctx.FormFile("photo")
	if errFile != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid read file"),
			infrafiber.WithError(errFile),
			infrafiber.WithHttpCode(http.StatusNotFound),
		).Send(ctx)
	}

	// create photo
	if err := h.svc.CreatePhoto(context.Background(), ctx, file, user_id); err != nil {
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
		infrafiber.WithMessage("upload photo success"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h handler) UploadDataTPS(ctx *fiber.Ctx) error {

	user_id := fmt.Sprintf("%+v", ctx.Locals("PUBLIC_ID"))

	file, _ := ctx.FormFile("photo")

	values := []string{"paslon1", "paslon2", "paslon3", "paslon4", "suara_sah", "suara_tidak_sah"}
	var result []string

	for _, v := range values {
		value := ctx.FormValue(v)
		if value == "" {
			return infrafiber.NewResponse(
				infrafiber.WithMessage("invalid read file for "+v),
				infrafiber.WithHttpCode(http.StatusNotFound),
			).Send(ctx)
		}
		result = append(result, value)
	}

	// create photo
	if err := h.svc.UploadDataTPS(context.Background(), ctx, result, file, user_id); err != nil {
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
		infrafiber.WithMessage("upload data success"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h handler) TPSAdressDetail(ctx *fiber.Ctx) error {

	user_id := fmt.Sprintf("%+v", ctx.Locals("PUBLIC_ID"))

	model, err := h.svc.TPSAdressDetail(context.Background(), user_id)
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

	tpsEntity := model.ToTpsDetailResponse()

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get tps success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(tpsEntity),
	).Send(ctx)
}

func (h handler) GetVoterTPS(ctx *fiber.Ctx) error {
	codeTPS := ctx.Params("code", "")
	if codeTPS == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	tps, err := h.svc.GetVoterTPS(context.Background(), codeTPS)
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

	getVoterTPSResponse := tps.ToGetVoterTPSResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get voter tps success"),
		infrafiber.WithPayload(getVoterTPSResponse),
	).Send(ctx)
}

func (h handler) GetAllVoterTPS(ctx *fiber.Ctx) error {

	tps, err := h.svc.GetAllVoterTPS(context.Background())

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

	getAllVoterTPSResponse := tps.ToGetAllVoterTPSResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get all voter tps success"),
		infrafiber.WithPayload(getAllVoterTPSResponse),
	).Send(ctx)
}

func (h handler) GetTPSSaksiPagination(ctx *fiber.Ctx) error {
	var req = ListTPSSaksiRequestPayload{}

	if err := ctx.QueryParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("list tps invalid"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	tpss, err := h.svc.ListTPSSaksi(context.Background(), req)
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

	ListTPSSaksiResponse := NewTPSListSaksiResponseFromEntity(tpss)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list tps success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(ListTPSSaksiResponse),
		infrafiber.WithQuery(req.GenerateDefaultValue()),
	).Send(ctx)
}

func (h handler) GetListTPS(ctx *fiber.Ctx) error {

	tpss, err := h.svc.GetListTPS(context.Background())
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

	kecamatanListResponse := NewTPSListResponseFromEntity(tpss)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list kelurahan success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(kecamatanListResponse),
	).Send(ctx)
}

func (h handler) UpdateVoteTPS(ctx *fiber.Ctx) error {
	var req = EditVoteTPSRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("list tps invalid"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	tpsId := ctx.Params("id", "")
	if tpsId == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	err := h.svc.EditVoteTPS(context.Background(), req, tpsId)
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

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("update data tps success"),
	).Send(ctx)
}

func (h handler) UpdateVoteTPSByUser(ctx *fiber.Ctx) error {
	var req = EditVoteTPSBySaksiRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("data payload invalid"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	userId := ctx.Params("id", "")
	if userId == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid user id"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	data, err := h.svc.UpdateVoteTPSByUserId(context.Background(), req, userId)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid data tps"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	dataPayload := data.ToTpsDetailFromUpdateDataResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage(fmt.Sprintf("Data Kecamatan %+s, Kelurahan %+s, %+s Berhasil ditambahkan", dataPayload.KecamatanName, dataPayload.KelurahanName, dataPayload.TpsName)),
	).Send(ctx)
}
