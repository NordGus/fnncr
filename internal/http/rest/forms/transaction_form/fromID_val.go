package transaction_form

import "github.com/google/uuid"

type FromID struct {
	Value  uuid.UUID `json:"value"`
	Errors []error   `json:"errors"`
}
