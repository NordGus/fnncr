package account_form

import "github.com/google/uuid"

type ParentID struct {
	Value  uuid.UUID `json:"value"`
	Errors []error   `json:"errors"`
}
