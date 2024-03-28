package transaction_entity

import (
	"time"

	account "financo/internal/entities/account_entity"
	"github.com/google/uuid"
)

type Entity struct {
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

	From account.Entity `json:"-"`
	To   account.Entity `json:"-"`
}

func New(
	id uuid.UUID,
	from account.Entity,
	to account.Entity,
	fromAmount int64,
	toAccount int64,
	issuedAt time.Time,
	executedAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) Entity {
	return Entity{
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
		From:       from,
		To:         to,
	}
}
