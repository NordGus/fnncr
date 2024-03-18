package transaction_form

import "time"

type IssuedAt struct {
	Value  time.Time
	Errors []error
}
