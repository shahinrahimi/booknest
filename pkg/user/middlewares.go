package user

import (
	"context"
	"net/http"
	"os/user"

	"github.com/shahinrahimi/booknest/types"
	"github.com/shahinrahimi/booknest/utils"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func (h *Handler) MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// deserilizing from json
		u := user.User{}
		if err := utils.FromJSON(&u, r.Body); err != nil {
			h.logger.Println("error deserializing user", err)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "bad request"})
			return
		}
		// validate user
		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(u); err != nil {
			h.logger.Println("validating user failed", err)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "bad request"})
			return
		}
		// Add the user to context
		ctx := context.WithValue(r.Context(), KeyUser{}, u)
		r = r.WithContext(ctx)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
