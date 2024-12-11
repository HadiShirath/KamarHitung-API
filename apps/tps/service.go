package tps

import (
	"context"
	"fmt"
	"kamar-hitung/apps/auth"
	"kamar-hitung/infra/response"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Repository interface {
	CreatePhoto(ctx context.Context, model TPS, userId string) (err error)
	UploadDataTPS(ctx context.Context, model TPS, userId string) (err error)
	GetAddressTPSByUserId(ctx context.Context, userId string) (tps TPS, err error)
	GetAllVoterTPS(ctx context.Context) (tps TPS, err error)
	GetVoterTPS(ctx context.Context, codeTPS string) (tps TPS, err error)
	GetAllTPSSaksiWithPaginationCursor(ctx context.Context, model TPSPagination) (tpss []TPS, err error)
	EditTPSSaksi(ctx context.Context, model TPS, userId string) (err error)
	GetListTPS(ctx context.Context) (tpss []TPS, err error)
	EditVoteTPS(ctx context.Context, model TPS, tpsId string) (err error)
	UpdateVoteTPSByUserId(ctx context.Context, model TPS, userId string) (updatedModel TPS, err error)
	GetListTPSCode(ctx context.Context, codeKelurahan string) (tpss []TPS, err error)
}

type service struct {
	repo     Repository
	repoAuth auth.Repository
}

func NewService(repo Repository, repoAuth auth.Repository) service {
	return service{
		repo:     repo,
		repoAuth: repoAuth,
	}
}

func (s service) GetListTPSCode(ctx context.Context, codeKelurahan string) (tpss []TPS, err error) {
	tpss, err = s.repo.GetListTPSCode(ctx, codeKelurahan)

	if err != nil {
		if err == response.ErrNotFound {
			return []TPS{}, nil
		}
		return
	}

	if len(tpss) == 0 {
		return []TPS{}, nil
	}
	return
}

func (s service) CreatePhoto(ctx context.Context, c *fiber.Ctx, file *multipart.FileHeader, userId string) (err error) {
	filenameFromUser := file.Filename
	ext_photo := filepath.Ext(filenameFromUser)

	dataUser, err := s.repo.GetAddressTPSByUserId(ctx, userId)
	if err != nil {
		return
	}

	// jika photo sebelumnya ada, maka hapus foto sebelumnya
	if dataUser.Photo != "" {
		filePath := filepath.Join("./public/images/", dataUser.Photo)
		os.Remove(filePath)
	}

	// nama file photo untuk simpan database
	timestamp := time.Now().Unix()
	filenameToDB := fmt.Sprintf("%+s-%+s-%+s-%+v%+s", dataUser.KecamatanName, dataUser.KelurahanName, dataUser.TpsName, timestamp, ext_photo)

	// simpan file di direktori
	errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/images/%s", filenameToDB))
	if errSaveFile != nil {
		return errSaveFile
	}

	productEntity := NewFromCreatePhotoRequest(
		CreatePhotoRequestPayload{
			Photo: filenameToDB,
		})

	if err = s.repo.CreatePhoto(ctx, productEntity, userId); err != nil {
		return
	}

	return

}

func (s service) UploadDataTPS(ctx context.Context, c *fiber.Ctx, values []string, file *multipart.FileHeader, userId string) (err error) {

	tpsEntity := NewFromUploadDataRequest(
		UploadDataRequestPayload{
			Paslon1:       values[0],
			Paslon2:       values[1],
			Paslon3:       values[2],
			Paslon4:       values[3],
			SuaraSah:      values[4],
			SuaraTidakSah: values[5],
		})

	err = tpsEntity.ValidateSuaraSah()
	if err != nil {
		return err
	}

	dataUser, err := s.repo.GetAddressTPSByUserId(ctx, userId)
	if err != nil {
		return err
	}

	if file != nil {
		filenameFromUser := file.Filename
		ext_photo := filepath.Ext(filenameFromUser)

		// jika photo sebelumnya ada, maka hapus foto sebelumnya
		if dataUser.Photo != "" {
			filePath := filepath.Join("./public/images/", dataUser.Photo)
			os.Remove(filePath)
		}

		// nama file photo untuk simpan database
		timestamp := time.Now().Unix()
		filenameToDB := fmt.Sprintf("%+s-%+s-%+s-%+v%+s", dataUser.KecamatanName, dataUser.KelurahanName, dataUser.TpsName, timestamp, ext_photo)

		// simpan file di direktori
		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/images/%s", filenameToDB))
		if errSaveFile != nil {
			return errSaveFile
		}

		tpsEntity.Photo = filenameToDB
	}

	if err = s.repo.UploadDataTPS(ctx, tpsEntity, dataUser.UserId); err != nil {
		return
	}

	return

}

func (s service) CheckPhoto(ctx context.Context, userId string) (photo string, err error) {
	model, err := s.repo.GetAddressTPSByUserId(ctx, userId)
	if err != nil {
		return
	}

	return model.Photo, nil
}

func (s service) TPSAdressDetail(ctx context.Context, userId string) (tps TPS, err error) {

	model, err := s.repo.GetAddressTPSByUserId(ctx, userId)
	if err != nil {
		return
	}

	return model, nil
}

func (s service) GetAllVoterTPS(ctx context.Context) (tps TPS, err error) {
	model, err := s.repo.GetAllVoterTPS(ctx)

	if err != nil {
		return
	}

	return model, nil
}

func (s service) GetVoterTPS(ctx context.Context, codeTPS string) (tps TPS, err error) {

	model, err := s.repo.GetVoterTPS(ctx, codeTPS)
	if err != nil {
		return
	}

	return model, nil
}

func (s service) ListTPSSaksi(ctx context.Context, req ListTPSSaksiRequestPayload) (tpss []TPS, err error) {
	pagination := NewTPSSaksiPaginationFromProductRequest(req)

	tpss, err = s.repo.GetAllTPSSaksiWithPaginationCursor(ctx, pagination)
	if err != nil {
		if err == response.ErrNotFound {
			return []TPS{}, nil
		}
		return
	}

	if len(tpss) == 0 {
		return []TPS{}, nil
	}

	return
}

func (s service) EditTPSSaksi(ctx context.Context, req EditTPSSaksiRequestPayload, userId string) (err error) {
	saksiEntity := NewFromEditTPSSaksiRequest(req)

	if err = s.repo.EditTPSSaksi(ctx, saksiEntity, userId); err != nil {
		return
	}

	return

}

func (s service) GetListTPS(ctx context.Context) (tpss []TPS, err error) {

	tpss, err = s.repo.GetListTPS(ctx)

	if err != nil {
		if err == response.ErrNotFound {
			return []TPS{}, nil
		}
		return
	}

	if len(tpss) == 0 {
		return []TPS{}, nil
	}
	return
}

func (s service) EditVoteTPS(ctx context.Context, req EditVoteTPSRequestPayload, tpsId string) (err error) {
	voteTPSEntity := NewFromEditVoteTPSRequest(req)

	if err = s.repo.EditVoteTPS(ctx, voteTPSEntity, tpsId); err != nil {
		return
	}

	return
}

func (s service) UpdateVoteTPSByUserId(ctx context.Context, req EditVoteTPSBySaksiRequestPayload, username string) (tps TPS, err error) {

	voteTPSEntity := NewFromEditVoteBySaksiTPSRequest(req)

	if err = validateCodeUnique(req.CodeUnique); err != nil {
		return
	}

	model, err := s.repoAuth.GetAuthByUsername(ctx, username)
	if err != nil {
		return
	}

	tps, err = s.repo.UpdateVoteTPSByUserId(ctx, voteTPSEntity, model.PublicId.String())
	if err != nil {
		return
	}

	return
}
