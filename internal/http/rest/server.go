package rest

// TODO: migrate primary/web/api here

type (
	Service struct {
	}

	Option func(s *Service)
)

func New(opts ...Option) *Service {
	s := &Service{}

	for i := 0; i < len(opts); i++ {
		opts[i](s)
	}

	return s
}

func (s *Service) Start() error {
	// TODO: migrate code from previous version
	panic("needs to be implemented")
}
