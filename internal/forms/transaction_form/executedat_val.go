package transaction_form

import "time"

type ExecutedAt struct {
	Value  time.Time `json:"value"`
	Errors []error   `json:"errors"`
}
