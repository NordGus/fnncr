package nullable_value

type (
	Validator[Val any] func(value Val) error

	Value[Val any] struct {
		Value  Val     `json:"value"`
		Errors []error `json:"errors"`

		validators []Validator[Val]
		validated  bool
		present    bool
	}
)

func New[Val any](value Val, validators ...Validator[Val]) Value[Val] {
	return Value[Val]{
		Value:      value,
		Errors:     make([]error, 0, 5),
		validators: validators,
		present:    true,
	}
}

func (v *Value[Val]) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Value[Val]) Validate() {
	if v.validated {
		return
	}

	for i := 0; i < len(v.validators); i++ {
		if err := v.validators[i](v.Value); err != nil {
			v.Errors = append(v.Errors, err)
		}
	}

	v.validated = true
}
