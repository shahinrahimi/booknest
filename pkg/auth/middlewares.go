package auth

import (
	"net/http"

	"github.com/shahinrahimi/booknest/types"
	"github.com/shahinrahimi/booknest/utils"
)

func (h *Handler) MiddlewareRequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		session, _ := h.cs.Get(r, "session")
		// Authentication
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			utils.WriteJSON(rw, http.StatusUnauthorized, types.ApiError{Error: "Authentication is required and has failed or has not yet been provided."})
			return
		}
		// Authorization
		if isAdmin, ok := session.Values["is_admin"].(bool); !ok || !isAdmin {
			utils.WriteJSON(rw, http.StatusForbidden, types.ApiError{Error: "You do not have permission to access this resource."})
			return
		}
		next.ServeHTTP(rw, r)
	})
}
