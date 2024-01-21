package web

import (
	"net/http"
	"time"
)

type AppOption func(a *App)

func DefaultAppOptions(a *App) {
	a.Server = &http.Server{
		Addr:        ":4269",
		ReadTimeout: 10 * time.Second,
	}
}
