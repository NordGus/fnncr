package account_form

import (
	model "financo/internal/http/rest/models/shared"
)

type Currency struct {
	Value  model.Currency `json:"value"`
	Errors []error        `json:"errors"`
}
