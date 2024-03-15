package signout

type Response struct {
	err error
}

func (resp *Response) Error() error {
	return resp.err
}
