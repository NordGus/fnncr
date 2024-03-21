package transaction_form

import (
	"time"

	"financo/internal/forms/shared/form_value"
)

type (
	Transaction interface {
		GetID() string
		GetFromID() string
		GetToID() string

		GetFromAmount() int64
		GetToAmount() int64

		GetIssuedAt() time.Time
		GetExecutedAt() time.Time
	}

	Form struct {
		ID     form_value.Value[string] `json:"id"`
		FromID form_value.Value[string] `json:"fromID"`
		ToID   form_value.Value[string] `json:"toID"`

		FromAmount form_value.Value[int64] `json:"fromAmount"`
		ToAmount   form_value.Value[int64] `json:"toAmount"`

		IssuedAt   form_value.Value[time.Time] `json:"issuedAt"`
		ExecutedAt form_value.Value[time.Time] `json:"executedAt"`

		IsValid bool `json:"isValid"`
	}
)

func NewEntry(raw Form, validator Validator) Form {
	f := Form{
		ID:         form_value.New(raw.ID.Value, validator.IDValidators()...),
		FromID:     form_value.New(raw.FromID.Value, validator.FromIDValidators()...),
		ToID:       form_value.New(raw.ToID.Value, validator.ToIDValidators()...),
		FromAmount: form_value.New(raw.FromAmount.Value, validator.FromAmountValidators()...),
		ToAmount:   form_value.New(raw.ToAmount.Value, validator.ToAmountValidators()...),
		IssuedAt:   form_value.New(raw.IssuedAt.Value, validator.IssuedAtValidators()...),
		ExecutedAt: form_value.New(raw.ExecutedAt.Value, validator.IssuedAtValidators()...),
	}

	f.Validate()

	return f
}

func New(t Transaction) Form {
	f := Form{
		ID:         form_value.New(t.GetID()),
		FromID:     form_value.New(t.GetFromID()),
		ToID:       form_value.New(t.GetToID()),
		FromAmount: form_value.New(t.GetFromAmount()),
		ToAmount:   form_value.New(t.GetToAmount()),
		IssuedAt:   form_value.New(t.GetIssuedAt()),
		ExecutedAt: form_value.New(t.GetExecutedAt()),
		IsValid:    true,
	}

	return f
}

func (f *Form) Validate() {
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
}
