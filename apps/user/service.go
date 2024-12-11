package user

import (
	"context"
	"kamar-hitung/infra/response"
	"kamar-hitung/internal/config"
)

type Repository interface {
	GetUserList(ctx context.Context) (users []User, err error)
	GetUserSaksiList(ctx context.Context) (users []User, err error)
	EditTPSSaksi(ctx context.Context, model User, userId string) (err error)
	GetDataForExportCSV(ctx context.Context) (users []User, err error)
	DeleteUserById(ctx context.Context, userId string) (err error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) GetUserList(ctx context.Context) (users []User, err error) {
	users, err = s.repo.GetUserList(ctx)
	if err != nil {
		if err == response.ErrNotFound {
			return []User{}, err
		}
		return
	}

	if len(users) == 0 {
		return []User{}, nil
	}

	return
}

func (s service) GetUserSaksiList(ctx context.Context) (users []User, err error) {
	users, err = s.repo.GetUserSaksiList(ctx)
	if err != nil {
		if err == response.ErrNotFound {
			return []User{}, err
		}
		return
	}

	if len(users) == 0 {
		return []User{}, nil
	}

	return
}

func (s service) EditTPSSaksi(ctx context.Context, req EditUserRequestPayload, userId string) (err error) {
	saksiEntity := NewFromEditTPSSaksiRequest(req)

	if saksiEntity.Password != "" {
		if err = saksiEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
			return
		}
	}

	if err = s.repo.EditTPSSaksi(ctx, saksiEntity, userId); err != nil {
		return
	}

	return

}

func (s service) GetDataForExportCSV(ctx context.Context) (data []User, err error) {

	data, err = s.repo.GetDataForExportCSV(ctx)

	if err != nil {
		if err == response.ErrNotFound {
			return []User{}, nil
		}
		return
	}

	if len(data) == 0 {
		return []User{}, nil
	}
	return
}

func (s service) DeleteUserById(ctx context.Context, userId string) (err error) {

	if err := s.repo.DeleteUserById(ctx, userId); err != nil {
		return err
	}

	return
}
