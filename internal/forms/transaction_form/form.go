package transaction_form

import (
	"time"

	"financo/internal/forms/shared/form_value"
)

type (
	Transaction interface {
		ID() string
		FromID() string
		ToID() string

		FromAmount() int64
		ToAmount() int64

		IssuedAt() time.Time
		ExecutedAt() time.Time
	}

	Form struct {
		ID     form_value.Value[string] `json:"id"`
		FromID form_value.Value[string] `json:"fromID"`
		ToID   form_value.Value[string] `json:"toID"`

		FromAmount form_value.Value[int64] `json:"fromAmount"`
		ToAmount   form_value.Value[int64] `json:"toAmount"`

		IssuedAt   form_value.Value[time.Time] `json:"issuedAt"`
		ExecutedAt form_value.Value[time.Time] `json:"executedAt"`

		IsValid     bool `json:"isValid"`
		initialized bool
	}
)

func NewEntry(raw Form, validator Validator) Form {
	f := Form{
		ID:          form_value.New(raw.ID.Value, validator.IDValidators()...),
		FromID:      form_value.New(raw.FromID.Value, validator.FromIDValidators()...),
		ToID:        form_value.New(raw.ToID.Value, validator.ToIDValidators()...),
		FromAmount:  form_value.New(raw.FromAmount.Value, validator.FromAmountValidators()...),
		ToAmount:    form_value.New(raw.ToAmount.Value, validator.ToAmountValidators()...),
		IssuedAt:    form_value.New(raw.IssuedAt.Value, validator.IssuedAtValidators()...),
		ExecutedAt:  form_value.New(raw.ExecutedAt.Value, validator.IssuedAtValidators()...),
		initialized: true,
	}

	f.Valid()

	return f
}

func New(transaction Transaction) Form {
	f := Form{
		ID:          form_value.New(transaction.ID()),
		FromID:      form_value.New(transaction.FromID()),
		ToID:        form_value.New(transaction.ToID()),
		FromAmount:  form_value.New(transaction.FromAmount()),
		ToAmount:    form_value.New(transaction.ToAmount()),
		IssuedAt:    form_value.New(transaction.IssuedAt()),
		ExecutedAt:  form_value.New(transaction.ExecutedAt()),
		initialized: true,
	}

	f.Valid()

	return f
}

func (f *Form) Initialized() bool {
	return f.initialized
}

func (f *Form) Valid() bool {
	f.ID.Validate()
	f.FromAmount.Validate()
	f.ToID.Validate()
	f.FromAmount.Validate()
	f.ToAmount.Validate()
	f.IssuedAt.Validate()
	f.ExecutedAt.Validate()

	f.IsValid = f.ID.Valid() &&
		f.FromAmount.Valid() &&
		f.ToID.Valid() &&
		f.FromAmount.Valid() &&
		f.ToAmount.Valid() &&
		f.IssuedAt.Valid() &&
		f.ExecutedAt.Valid()

	return f.IsValid
}
