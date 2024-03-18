package account_form

import (
	"financo/internal/entities/shared"
)

type Currency struct {
	Value  shared.Currency `json:"value"`
	Errors []error         `json:"errors"`
}
