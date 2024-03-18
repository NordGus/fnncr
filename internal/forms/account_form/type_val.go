package account_form

import (
	model "financo/internal/entities/account_entity"
)

type Type struct {
	Value  model.Type `json:"value"`
	Errors []error    `json:"errors"`
}
