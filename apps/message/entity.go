package message

import (
	"time"

	"github.com/google/uuid"
)

type Inbox struct {
	Id           string    `db:"id"`
	SenderNumber string    `db:"sender_number"`
	Message      string    `db:"message"`
	UpdatedAt    time.Time `db:"updated_at"`
	CreatedAt    string    `db:"created_at"`
}

type Outbox struct {
	Id              string    `db:"id"`
	ReceiverNumber  string    `db:"receiver_number"`
	ReceiverNumbers []string  `db:"receiver_numbers"`
	Message         string    `db:"message"`
	Processed       bool      `db:"processed"`
	UpdatedAt       time.Time `db:"updated_at"`
	CreatedAt       string    `db:"created_at"`
}

type StatusMessage struct {
	Processed string `json:"processed"`
}

func NewFromCreateMessageRequest(req CreateMessageRequestPayload) Outbox {
	return Outbox{
		Id:             uuid.NewString(),
		ReceiverNumber: req.ReceiverNumber,
		Message:        req.Message,
		Processed:      req.Processed,
	}
}

func NewFromCreateMessagesRequest(req CreateMessagesRequestPayload) Outbox {
	return Outbox{
		Id:              uuid.NewString(),
		ReceiverNumbers: req.ReceiverNumbers,
		Message:         req.Message,
		Processed:       req.Processed,
	}
}

func NewFromUploadInboxRequest(req UploadInboxRequestPayload) Inbox {
	return Inbox{
		Id:           uuid.NewString(),
		SenderNumber: req.SenderNumber,
		Message:      req.Message,
	}
}

func NewFromStatusMessageRequest(req StatusMessageRequestPayload) StatusMessage {
	return StatusMessage{
		Processed: req.Processed,
	}
}

func (i Inbox) ToInboxListResponse() InboxListResponse {
	return InboxListResponse{
		Id:           i.Id,
		SenderNumber: i.SenderNumber,
		Message:      i.Message,
		CreatedAt:    i.CreatedAt,
		UpdatedAt:    i.UpdatedAt,
	}
}

func (o Outbox) ToOutboxListResponse() OutboxListResponse {
	return OutboxListResponse{
		Id:              o.Id,
		ReceiverNumber:  o.ReceiverNumber,
		ReceiverNumbers: o.ReceiverNumbers,
		Message:         o.Message,
		Processed:       o.Processed,
		CreatedAt:       o.CreatedAt,
		UpdatedAt:       o.UpdatedAt,
	}
}

func ConvertTimestamps(timeCurrent string) (formattedTime string) {
	// Parse timestamp ke dalam waktu
	t, err := time.Parse(time.RFC3339Nano, timeCurrent)
	if err != nil {
		return timeCurrent
	}

	// Konversi waktu ke WIB (UTC+7)
	wib := t.In(time.FixedZone("WIB", 7*60*60))

	// Format ketentuan Go dalam menentukan waktu
	formattedTime = wib.Format("2006-01-02, 15:04") + " WIB"
	return formattedTime
}
