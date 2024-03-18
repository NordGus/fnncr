package transaction_model

import (
	"time"

	account "financo/internal/http/rest/models/account_model"
	"github.com/google/uuid"
)

type Model struct {
	ID     uuid.UUID
	FromID uuid.UUID
	ToID   uuid.UUID

	FromAmount int64
	ToAmount   int64

	IssuedAt   time.Time
	ExecutedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func New(
	id uuid.UUID,
	from account.Model,
	to account.Model,
	fromAmount int64,
	toAccount int64,
	issuedAt time.Time,
	executedAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) Model {
	return Model{
		ID:         id,
		FromID:     from.ID,
		ToID:       to.ID,
		FromAmount: fromAmount,
		ToAmount:   toAccount,
		IssuedAt:   issuedAt,
		ExecutedAt: executedAt,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  deletedAt,
	}
}
