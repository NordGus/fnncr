package account_form

import (
	account "financo/internal/entities/account_entity"
)

type Kind struct {
	Value  account.Kind `json:"value"`
	Errors []error      `json:"errors"`
}
