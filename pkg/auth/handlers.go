package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
	contentType := r.Header.Get("Content-Type")
	switch {
	// handle JSON request
	case strings.HasPrefix(contentType, "application/json"):
		if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
			h.logger.Println("error decerelizing login request", err)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "bad request"})
			return
		}
		// handle formdata request => htmx
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		if err := r.ParseForm(); err != nil {
			h.logger.Println("error parsing form data", err)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "bad request"})
			return
		}
		loginRequest.Username = r.FormValue("username")
		loginRequest.Password = r.FormValue("password")
		h.logger.Printf("username: %s password: %s", loginRequest.Username, loginRequest.Password)
	default:
		h.logger.Println("unsupported content type", contentType)
		utils.WriteJSON(rw, http.StatusUnsupportedMediaType, types.ApiError{Error: "unsupported content type"})
		return
	}
	// check if username is exists
	u, err := h.store.GetUser(loginRequest.Username)
	if err != nil {
		h.logger.Println("error query user", err)
		utils.WriteJSON(rw, http.StatusUnauthorized, types.ApiError{Error: "username or password is wrong"})
		return
	}

	// check password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginRequest.Password)); err != nil {
		h.logger.Println("error checking passwords", err)
		utils.WriteJSON(rw, http.StatusUnauthorized, types.ApiError{Error: "username or password is wrong"})
		return
	}

	// set session for user
	session, _ := h.cs.Get(r, Session)
	session.Values[IsAuthenticated] = true
	session.Values[IsAdmin] = u.IsAdmin
	session.Values[Username] = u.Username
	session.Values[UserId] = u.ID
	if err := session.Save(r, rw); err != nil {
		h.logger.Println("error saving session for user", err)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}
	// set custome htmx header
	rw.Header().Set("HX-Redirect", "/")
	utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: "logged in successfully"})
}

func (h *Handler) Logout(rw http.ResponseWriter, r *http.Request) {
	session, _ := h.cs.Get(r, Session)
	session.Values[IsAuthenticated] = false
	session.Values[IsAdmin] = false
	session.Values[Username] = ""
	session.Values[UserId] = ""
	if err := session.Save(r, rw); err != nil {
		h.logger.Println("error saving session for user", err)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}
	// set custome htmx header
	rw.Header().Set("HX-Redirect", "/")
	rw.Header().Set("HX-Refresh", "true")
	// write response
	utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: "logged out successfully"})
}
