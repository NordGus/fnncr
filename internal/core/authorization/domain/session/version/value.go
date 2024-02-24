package version

type Value uint32

func New(version uint32) (Value, error) {
	return Value(version), nil
}

func (v Value) Uint32() uint32 {
	return uint32(v)
}

func (v Value) IsInvalid(version uint32) bool {
	return v.Uint32() != version
}
