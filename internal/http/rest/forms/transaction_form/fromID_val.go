package transaction_form

import "github.com/google/uuid"

type FromID struct {
	Value  uuid.UUID
	Errors []error
}
