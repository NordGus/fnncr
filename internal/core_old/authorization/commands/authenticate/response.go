package authenticate

import "financo/internal/core_old/authorization/domain/user"

type Response struct {
	user user.Entity
	err  error
}

func (resp *Response) User() user.Entity {
	return resp.user
}

func (resp *Response) Error() error {
	return resp.err
}
