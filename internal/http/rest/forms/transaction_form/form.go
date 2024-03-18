package transaction_form

import (
	"time"

	"github.com/google/uuid"
)

type Form struct {
	FromID FromID `json:"fromID"`
	ToID   ToID   `json:"toID"`

	FromAmount FromAmount `json:"fromAmount"`
	ToAmount   ToAmount   `json:"toAmount"`

	IssuedAt   IssuedAt   `json:"issuedAt"`
	ExecutedAt ExecutedAt `json:"executedAt"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}

func New(
	from uuid.UUID,
	to uuid.UUID,
	fromAmount int64,
	toAccount int64,
	issuedAt time.Time,
	executedAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) Form {
	f := Form{
		FromID:     FromID{Value: from, Errors: nil},
		ToID:       ToID{Value: to, Errors: nil},
		FromAmount: FromAmount{Value: fromAmount, Errors: nil},
		ToAmount:   ToAmount{Value: toAccount, Errors: nil},
		IssuedAt:   IssuedAt{Value: issuedAt, Errors: nil},
		ExecutedAt: ExecutedAt{Value: executedAt, Errors: nil},
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  deletedAt,
	}

	return f
}
