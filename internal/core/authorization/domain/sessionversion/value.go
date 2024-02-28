package sessionversion

type Value struct {
	value uint32
}

func New(version uint32) (Value, error) {
	return Value{value: version}, nil
}

func (v Value) Uint32() uint32 {
	return v.value
}

func (v Value) IsInvalid(current Value) bool {
	return v.value != current.value
}
