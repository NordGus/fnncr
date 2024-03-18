package account_form

type IsArchive struct {
	Value  bool    `json:"value"`
	Errors []error `json:"errors"`
}
