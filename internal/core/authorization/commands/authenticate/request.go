package authenticate

import "context"

type Request struct {
	ctx       context.Context
	sessionID string
}

func NewRequest(ctx context.Context, sessionID string) Request {
	return Request{
		ctx:       ctx,
		sessionID: sessionID,
	}
}
