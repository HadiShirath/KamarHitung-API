package message

import (
	"context"
	"database/sql"
	"kamar-hitung/infra/response"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) GetInboxList(ctx context.Context) (inboxs []Inbox, err error) {
	query := `
	SELECT
		id, sender_number, message, created_at, updated_at
	FROM 
		inbox
	ORDER BY
		created_at DESC
	`

	err = r.db.SelectContext(ctx, &inboxs, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}

	return

}

func (r repository) GetOutboxList(ctx context.Context, model StatusMessage) (outboxs []Outbox, err error) {

	var query string

	if model.Processed != "" {
		query = `
		SELECT
			id, receiver_number, receiver_numbers, message, processed, created_at, updated_at
		FROM 
			outbox
		WHERE
			processed=$1
		ORDER BY
			created_at ASC
		`
		rows, err := r.db.QueryContext(ctx, query, model.Processed)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, response.ErrNotFound
			}

		}

		defer rows.Close()

		for rows.Next() {
			var outbox Outbox
			if err := rows.Scan(&outbox.Id, &outbox.ReceiverNumber, pq.Array(&outbox.ReceiverNumbers), &outbox.Message, &outbox.Processed, &outbox.CreatedAt, &outbox.UpdatedAt); err != nil {
				log.Fatalln(err)
			}
			outboxs = append(outboxs, outbox)
		}

	} else {
		query = `
			SELECT
				id, receiver_number, receiver_numbers, message, processed, created_at, updated_at
			FROM 
				outbox
			ORDER BY
				created_at DESC
			`

		rows, err := r.db.QueryContext(ctx, query)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, response.ErrNotFound
			}

		}

		defer rows.Close()

		for rows.Next() {
			var outbox Outbox
			if err := rows.Scan(&outbox.Id, &outbox.ReceiverNumber, pq.Array(&outbox.ReceiverNumbers), &outbox.Message, &outbox.Processed, &outbox.CreatedAt, &outbox.UpdatedAt); err != nil {
				log.Fatalln(err)
			}
			outboxs = append(outboxs, outbox)
		}
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}

	return

}

func (r repository) CreateMessage(ctx context.Context, model Outbox) (err error) {
	query := `
		INSERT INTO outbox (
			id, receiver_number, message, processed
		) VALUES (
			:id, :receiver_number, :message, :processed
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return
}

func (r repository) CreateMessages(ctx context.Context, model Outbox) (err error) {
	query := `
		INSERT INTO outbox (
			id, receiver_numbers, message, processed
		) VALUES (
			:id, :receiver_numbers, :message, :processed
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	// Menggunakan pq.Array saat mengeksekusi query
	if _, err = stmt.ExecContext(ctx, map[string]interface{}{
		"id":               model.Id,
		"receiver_numbers": pq.Array(model.ReceiverNumbers),
		"message":          model.Message,
		"processed":        model.Processed,
	}); err != nil {
		return
	}
	return
}

func (r repository) UploadInbox(ctx context.Context, model Inbox) (err error) {
	query := `
		INSERT INTO inbox (
			id, sender_number, message
		) VALUES (
			:id, :sender_number, :message
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return
}

func (r repository) GetOutboxById(ctx context.Context, outboxId string) (model Outbox, err error) {
	query := `
		SELECT
			id, receiver_number, message
		FROM outbox
		WHERE id=$1
	`

	err = r.db.GetContext(ctx, &model, query, outboxId)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) UpdateStatusOutbox(ctx context.Context, outboxId string) (err error) {
	query := `
		UPDATE outbox SET processed=true WHERE id=$1
	`

	_, err = r.db.ExecContext(ctx, query, outboxId)

	if err != nil {
		return
	}

	return
}
