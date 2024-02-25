package sessionversion

type Value struct {
	version uint32
}

func New(version uint32) (Value, error) {
	return Value{
		version: version,
	}, nil
}
