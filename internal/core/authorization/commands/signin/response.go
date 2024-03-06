package signin

type Response struct {
	sessionID string
	err       error
}

func (resp *Response) SessionID() string {
	return resp.sessionID
}

func (resp *Response) Error() error {
	return resp.err
}
