package authentication

type Service struct {
	sessionCookieName string
}

func New() *Service {
	return &Service{
		sessionCookieName: "_fnncr_session",
	}
}
