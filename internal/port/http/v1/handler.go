package v1

import (
	"docker_ex5/internal/app"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type handler struct {
	app app.Application
}

func NewHandler(app app.Application, router chi.Router) http.Handler {
	return HandlerFromMux(handler{
		app: app,
	}, router)
}
