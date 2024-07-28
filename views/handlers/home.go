package handlers

import (
	"net/http"

	"github.com/shahinrahimi/booknest/pkg/user"

	"github.com/shahinrahimi/booknest/views/home"
)

func (vh *ViewHandler) HandleHome(rw http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(user.KeyUser{}).(user.User)
	if err := render(rw, r, home.Home(u)); err != nil {
		vh.logger.Printf("Error handle home view: %v", err)
	}
}
