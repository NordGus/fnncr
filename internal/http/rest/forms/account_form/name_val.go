package account_form

type Name struct {
	Value  string  `json:"value"`
	Errors []error `json:"errors"`
}
