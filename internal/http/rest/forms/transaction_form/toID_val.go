package transaction_form

import "github.com/google/uuid"

type ToID struct {
	Value  uuid.UUID `json:"value"`
	Errors []error   `json:"errors"`
}
