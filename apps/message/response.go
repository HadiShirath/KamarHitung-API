package message

import (
	"time"
)

type InboxListResponse struct {
	Id           string    `json:"id"`
	SenderNumber string    `json:"sender_number"`
	Message      string    `json:"message"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    string    `json:"created_at"`
}
type OutboxListResponse struct {
	Id              string    `json:"id"`
	ReceiverNumber  string    `json:"receiver_number"`
	ReceiverNumbers []string  `json:"receiver_numbers"`
	Message         string    `json:"message"`
	Processed       bool      `json:"processed"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       string    `json:"created_at"`
}

func NewListInboxResponseFromEntity(inboxs []Inbox) []InboxListResponse {
	var ListInbox = []InboxListResponse{}

	for _, inbox := range inboxs {

		inbox.CreatedAt = ConvertTimestamps(inbox.CreatedAt)
		ListInbox = append(ListInbox, inbox.ToInboxListResponse())
	}

	return ListInbox
}

func NewListOutboxResponseFromEntity(outboxs []Outbox) []OutboxListResponse {
	var ListOutbox = []OutboxListResponse{}

	for _, outbox := range outboxs {

		outbox.CreatedAt = ConvertTimestamps(outbox.CreatedAt)
		ListOutbox = append(ListOutbox, outbox.ToOutboxListResponse())
	}

	return ListOutbox
}
