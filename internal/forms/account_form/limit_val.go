package account_form

type Limit struct {
	Value  int64   `json:"value"`
	Errors []error `json:"errors"`
}
