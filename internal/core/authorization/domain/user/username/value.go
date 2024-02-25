package username

type Value struct {
	username string
}

func New(username string) (Value, error) {
	return Value{
		username: username,
	}, nil
}
