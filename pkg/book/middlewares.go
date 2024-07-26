package book

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shahinrahimi/booknest/types"
	"github.com/shahinrahimi/booknest/utils"
)

var validate *validator.Validate

func (h *Handler) MiddlewareValidateBook(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// deserilizing from json
		var b Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			h.logger.Println("error deserializing user", err)
			utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
			return
		}

		// validate book
		validate = validator.New(validator.WithRequiredStructEnabled())

		if err := validate.Struct(b); err != nil {
			h.logger.Println("validating book failed", err)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "bad request"})
			return
		}
		// Add the book to the context
		ctx := context.WithValue(r.Context(), KeyBook{}, b)
		r = r.WithContext(ctx)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
