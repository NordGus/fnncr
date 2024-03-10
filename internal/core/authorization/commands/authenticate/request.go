package authenticate

import (
	"context"
	"financo/internal/core/authorization/domain/sessionID"
)

type Request struct {
	ctx       context.Context
	sessionID sessionID.Value
	err       error
}

func NewRequest(ctx context.Context, sessionId string, encoder sessionID.Encoder) Request {
	sid, err := sessionID.NewFromString(sessionId, encoder)

	return Request{
		ctx:       ctx,
		sessionID: sid,
		err:       err,
	}
}
