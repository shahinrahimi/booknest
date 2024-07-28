package handlers

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
)

type ViewHandler struct {
	logger *log.Logger
}

func NewViewHandler(logger *log.Logger) *ViewHandler {
	return &ViewHandler{logger}
}

func render(rw http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), rw)
}
