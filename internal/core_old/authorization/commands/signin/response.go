package signin

import "financo/internal/core_old/authorization/domain/sessionID"

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
