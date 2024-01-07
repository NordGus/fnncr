package authentication

import "context"

type Service struct {
	ctx       context.Context
	cancelCtx context.CancelCauseFunc
}

func New(ctx context.Context, cancelCtx context.CancelCauseFunc) *Service {
	return &Service{
		ctx:       ctx,
		cancelCtx: cancelCtx,
	}
}
