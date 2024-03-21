package account_form

import (
	"time"

	"financo/internal/forms/shared/form_value"
	"financo/internal/forms/shared/nullable_value"
)

type Validator interface {
	IDValidators() []nullable_value.Validator[string]
	ParentIDValidators() []nullable_value.Validator[string]

	KindValidators() []form_value.Validator[string]
	CurrencyValidators() []form_value.Validator[string]
	LimitValidators() []form_value.Validator[int64]
	IsArchiveValidators() []form_value.Validator[bool]

	NameValidators() []form_value.Validator[string]
	DescriptionValidators() []form_value.Validator[string]
	ColorValidators() []form_value.Validator[string]
	IconValidators() []form_value.Validator[string]

	InitialAmountValidators() []form_value.Validator[int64]
	OpenedAtValidators() []form_value.Validator[time.Time]
}
