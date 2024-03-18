package transaction_form

import (
	"time"

	"github.com/google/uuid"
)

type Form struct {
	FromID FromID
	ToID   ToID

	FromAmount FromAmount
	ToAmount   ToAmount

	IssuedAt   IssuedAt
	ExecutedAt ExecutedAt

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
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
