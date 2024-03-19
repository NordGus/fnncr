package account_form

import (
	"financo/internal/entities/shared/currencies"
)

type Currency struct {
	Value  currencies.Currency `json:"value"`
	Errors []error             `json:"errors"`
}
