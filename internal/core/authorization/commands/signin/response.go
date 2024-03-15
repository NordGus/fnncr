package signin

import "financo/internal/core/authorization/domain/sessionID"

type Response struct {
	sessionID sessionID.Value
	err       error
}

func (resp *Response) SessionID() sessionID.Value {
	return resp.sessionID
}

func (resp *Response) Error() error {
	return resp.err
}
