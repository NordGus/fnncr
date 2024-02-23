package api

import (
	"net/http"
	"time"
)

type AppOption func(a *App)

func DefaultAppOptions(a *App) {
	a.Server = &http.Server{
		Addr:        ":3000",
		Handler:     a.echo,
		ReadTimeout: 10 * time.Second,
	}
}
