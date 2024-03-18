package transaction_form

import "time"

type ExecutedAt struct {
	Value  time.Time
	Errors []error
}
