package signout

import (
	"context"
	"financo/internal/core_old/authorization/domain/user"
)

type Request struct {
	ctx  context.Context
	user user.Entity
}

func NewRequest(ctx context.Context, user user.Entity) Request {
	return Request{
		ctx:  ctx,
		user: user,
	}
}
