package message

import (
	"context"
	"kamar-hitung/apps/auth"
	"kamar-hitung/infra/response"
)

type Repository interface {
	GetInboxList(ctx context.Context) (inboxs []Inbox, err error)
	GetOutboxList(ctx context.Context, model StatusMessage) (outboxs []Outbox, err error)
	CreateMessage(ctx context.Context, model Outbox) (err error)
	CreateMessages(ctx context.Context, model Outbox) (err error)
	UploadInbox(ctx context.Context, model Inbox) (err error)
	GetOutboxById(ctx context.Context, outboxId string) (model Outbox, err error)
	UpdateStatusOutbox(ctx context.Context, outboxId string) (err error)
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

func (s service) GetInboxList(ctx context.Context) (inboxs []Inbox, err error) {

	inboxs, err = s.repo.GetInboxList(ctx)

	if err != nil {
		if err == response.ErrNotFound {
			return []Inbox{}, err
		}
		return
	}

	if len(inboxs) == 0 {
		return []Inbox{}, nil
	}

	return
}

func (s service) GetOutboxList(ctx context.Context, req StatusMessageRequestPayload) (outboxs []Outbox, err error) {

	status := NewFromStatusMessageRequest(req)

	if status.Processed != "true" && status.Processed != "false" {
		status.Processed = ""
	}

	outboxs, err = s.repo.GetOutboxList(ctx, status)

	if err != nil {
		if err == response.ErrNotFound {
			return []Outbox{}, err
		}
		return
	}

	if len(outboxs) == 0 {
		return []Outbox{}, nil
	}

	return
}

func (s service) CreateMessage(ctx context.Context, req CreateMessageRequestPayload) (err error) {
	messageEntity := NewFromCreateMessageRequest(req)

	if err = s.repo.CreateMessage(ctx, messageEntity); err != nil {
		return
	}

	return

}

func (s service) CreateMessages(ctx context.Context, req CreateMessagesRequestPayload) (err error) {
	messagesEntity := NewFromCreateMessagesRequest(req)

	if err = s.repo.CreateMessages(ctx, messagesEntity); err != nil {
		return
	}

	return

}

func (s service) UploadInbox(ctx context.Context, req UploadInboxRequestPayload) (err error) {
	messageEntity := NewFromUploadInboxRequest(req)

	if err = s.repo.UploadInbox(ctx, messageEntity); err != nil {
		return
	}

	return

}

func (s service) UpdateStatusOutbox(ctx context.Context, outboxId string) (err error) {

	outbox, err := s.repo.GetOutboxById(ctx, outboxId)
	if err != nil {
		return
	}

	if err = s.repo.UpdateStatusOutbox(ctx, outbox.Id); err != nil {
		return
	}

	return

}
