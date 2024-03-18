package transaction_form

type FromAmount struct {
	Value  int64   `json:"value"`
	Errors []error `json:"errors"`
}
