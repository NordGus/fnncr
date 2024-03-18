package transaction_model

import (
	"time"

	account "financo/internal/http/rest/models/account_model"
	"github.com/google/uuid"
)

type Model struct {
	ID     uuid.UUID `json:"id"`
	FromID uuid.UUID `json:"fromID"`
	ToID   uuid.UUID `json:"toID"`

	FromAmount int64 `json:"fromAmount"`
	ToAmount   int64 `json:"toAmount"`

	IssuedAt   time.Time `json:"issuedAt"`
	ExecutedAt time.Time `json:"executedAt"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
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
