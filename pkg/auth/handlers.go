package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/shahinrahimi/booknest/pkg/user"
	"github.com/shahinrahimi/booknest/types"
	"github.com/shahinrahimi/booknest/utils"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	logger *log.Logger
	store  user.Storage
	cs     *sessions.CookieStore
}

func NewHandler(logger *log.Logger, store user.Storage, cs *sessions.CookieStore) *Handler {
	return &Handler{logger, store, cs}
}

func (h *Handler) Login(rw http.ResponseWriter, r *http.Request) {
	var loginRequest types.LoginRequest
	// decode body and extract loginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		h.logger.Println("error decerelizing login request", err)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internral error"})
		return
	}
	// check if username is exists
	u, err := h.store.GetUser(loginRequest.Username)
	if err != nil {
		h.logger.Println("error query user", err)
		utils.WriteJSON(rw, http.StatusForbidden, types.ApiError{Error: "username or password is wrong"})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginRequest.Password)); err != nil {
		h.logger.Println("error checking passwords", err)
		utils.WriteJSON(rw, http.StatusForbidden, types.ApiError{Error: "username or password is wrong"})
		return
	}

	// set session for user
	session, _ := h.cs.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["username"] = u.Username
	session.Values["is_admin"] = u.IsAdmin
	if err := session.Save(r, rw); err != nil {
		h.logger.Println("error saving session for user", err)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: "logged in successfully"})
}

func (h *Handler) Logout(rw http.ResponseWriter, r *http.Request) {
	session, _ := h.cs.Get(r, "session")
	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Values["is_admin"] = false
	if err := session.Save(r, rw); err != nil {
		h.logger.Println("error saving session for user", err)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: "logged out successfully"})
}
