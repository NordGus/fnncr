package transaction_form

import (
	"time"

	"financo/internal/forms/shared/form_value"
)

type Validator interface {
	IDValidators() []form_value.Validator[string]
	FromIDValidators() []form_value.Validator[string]
	ToIDValidators() []form_value.Validator[string]

	FromAmountValidators() []form_value.Validator[int64]
	ToAmountValidators() []form_value.Validator[int64]

	IssuedAtValidators() []form_value.Validator[time.Time]
	ExecutedAtAtValidators() []form_value.Validator[time.Time]
}
