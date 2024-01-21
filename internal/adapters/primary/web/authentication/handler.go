package authentication

import (
	"github.com/NordGus/fnncr/internal/core/services/authentication"
)

type Handler struct {
	api authentication.API
}

func New(api authentication.API) *Handler {
	return &Handler{
		api: api,
	}
}
