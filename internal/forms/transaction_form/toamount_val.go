package transaction_form

type ToAmount struct {
	Value  int64   `json:"value"`
	Errors []error `json:"errors"`
}
