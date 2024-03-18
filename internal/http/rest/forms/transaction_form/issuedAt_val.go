package transaction_form

import "time"

type IssuedAt struct {
	Value  time.Time `json:"value"`
	Errors []error   `json:"errors"`
}
