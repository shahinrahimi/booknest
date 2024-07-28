package auth

import (
	"context"
	"net/http"

	"github.com/shahinrahimi/booknest/pkg/user"
	"github.com/shahinrahimi/booknest/types"
	"github.com/shahinrahimi/booknest/utils"
)

func (h *Handler) MiddlewareRequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		session, _ := h.cs.Get(r, Session)
		// Authentication
		if auth, ok := session.Values[IsAuthenticated].(bool); !ok || !auth {
			utils.WriteJSON(rw, http.StatusUnauthorized, types.ApiError{Error: "Authentication is required and has failed or has not yet been provided."})
			return
		}
		// Authorization
		if isAdmin, ok := session.Values[IsAdmin].(bool); !ok || !isAdmin {
			utils.WriteJSON(rw, http.StatusForbidden, types.ApiError{Error: "You do not have permission to access this resource."})
			return
		}
		next.ServeHTTP(rw, r)
	})
}

func (h *Handler) MiddlewareProvideAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var u user.User
		session, _ := h.cs.Get(r, Session)
		if auth, ok := session.Values[IsAuthenticated].(bool); !ok || !auth {
			u.Username = ""
			u.ID = ""
			u.IsAdmin = false
		} else {
			u.Username, _ = session.Values[Username].(string)
			u.ID, _ = session.Values[UserId].(string)
			u.IsAdmin = session.Values[IsAdmin].(bool)
		}
		ctx := context.WithValue(r.Context(), user.KeyUser{}, u)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)

	})
}
