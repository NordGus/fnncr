package account_form

import model "financo/internal/http/rest/models/account_model"

type Type struct {
	Value  model.Type
	Errors []error
}
