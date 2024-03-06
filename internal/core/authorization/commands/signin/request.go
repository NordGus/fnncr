package signin

import "context"

type Request struct {
	ctx      context.Context
	username string
	password string
}

func NewRequest(ctx context.Context, username string, password string) Request {
	return Request{
		ctx:      ctx,
		username: username,
		password: password,
	}
}
