package kecamatan

import (
	"context"
	"encoding/csv"
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

func (h handler) GetAllVoter(ctx *fiber.Ctx) error {

	kecamatan, err := h.svc.AllVoter(context.Background())
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

	getAllVoterResponse := kecamatan.ToAllVoterResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get voter kelurahan success"),
		infrafiber.WithPayload(getAllVoterResponse),
	).Send(ctx)

}

func (h handler) GetVoterKecamatan(ctx *fiber.Ctx) error {
	codeKecamatan := ctx.Params("code", "")
	if codeKecamatan == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	kecamatan, err := h.svc.GetVoterKecamatan(context.Background(), codeKecamatan)
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

	getVoterKecamatanResponse := kecamatan.ToGetVoterKecamatanResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get voter kelurahan success"),
		infrafiber.WithPayload(getVoterKecamatanResponse),
	).Send(ctx)

}

func (h handler) ListKecamatan(ctx *fiber.Ctx) error {

	kecamatans, err := h.svc.GetListKecamatan(context.Background())
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

	kecamatanListResponse := NewKecamatanListResponseFromEntity(kecamatans)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list kecamatan success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(kecamatanListResponse),
	).Send(ctx)
}

func (h handler) GetListKelurahanFromKecamatan(ctx *fiber.Ctx) error {
	codeKecamatan := ctx.Params("code", "")
	if codeKecamatan == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	kelurahans, err := h.svc.GetListKelurahanFromKecamatan(context.Background(), codeKecamatan)
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

	kecamatanListResponse := NewKelurahanListResponseFromEntity(kelurahans)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list kelurahan success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(kecamatanListResponse),
	).Send(ctx)
}

func (h handler) GetListKecamatanCode(ctx *fiber.Ctx) error {

	kecamatans, err := h.svc.GetListKecamatanCode(context.Background())
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

	kecamatanCodeListResponse := NewKecamatanCodeResponseFromEntity(kecamatans)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list kecamatan code success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(kecamatanCodeListResponse),
	).Send(ctx)
}

func (h handler) importCSVHandler(ctx *fiber.Ctx) error {

	file, err := ctx.FormFile("file")
	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("Error retrieving the file"),
			infrafiber.WithError(err),
		).Send(ctx)
	}

	// Membaca file CSV langsung dari input
	f, err := file.Open()
	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("Error opening the CSV file"),
			infrafiber.WithError(err),
		).Send(ctx)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("Error reading the CSV file"),
			infrafiber.WithError(err),
		).Send(ctx)
	}

	err = h.svc.ImportCSV(context.Background(), records)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid process import data"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("Data imported successfully!"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)

}
