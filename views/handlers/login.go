package handlers

import (
	"net/http"

	"github.com/shahinrahimi/booknest/views/auth"
)

func (vh *ViewHandler) HandlerLogin(rw http.ResponseWriter, r *http.Request) {
	render(rw, r, auth.Login())
}
