package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shahinrahimi/booknest/types"
	"github.com/shahinrahimi/booknest/utils"
)

type Handler struct {
	logger  *log.Logger
	storage Storage
}

func NewHandler(logger *log.Logger, storage Storage) *Handler {
	return &Handler{logger, storage}
}

func (h *Handler) ListAll(rw http.ResponseWriter, r *http.Request) {
	users, err := h.storage.GetUsers()
	if err != nil {
		h.logger.Println("error query user from DB", err)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}
	if len(users) == 0 {
		h.logger.Println("no user found in DB")
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiSuccess{Message: "there is no user found"})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, users)
}

func (h *Handler) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := h.storage.GetUser(id)
	if err != nil {
		h.logger.Println("error query user with id", id)
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the user id is not valid"})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, user)
}

func (h *Handler) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch validated user from the context
	u := r.Context().Value(KeyUser{}).(User)
	// check if username is not already taken by another user
	users, err := h.storage.GetUsers()
	if err != nil {
		h.logger.Println("error query user from DB", err)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}
	for _, fetchedUser := range users {
		if fetchedUser.Username == u.Username {
			h.logger.Println("duplicate username")
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "username taken"})
			return
		}
	}

	// create new user
	newUser := NewUser(u.Username, u.Password)
	if err := h.storage.CreateUser(*newUser); err != nil {
		h.logger.Println("error creating a new user")
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "failed to create a new user"})
		return
	}

	utils.WriteJSON(rw, http.StatusCreated, types.ApiSuccess{Message: fmt.Sprintf("new user created with id: %v", newUser.ID)})
}

func (h *Handler) Update(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// check if id is avaialable
	if _, err := h.storage.GetUser(id); err != nil {
		h.logger.Println("error query user with id", id)
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the user id is not valid"})
		return
	}
	// user := r.Context().Value(KeyUser{}).(User)
	utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: "the update method not implemented for user"})
}

func (h *Handler) Delete(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// check if id is avaialable
	if _, err := h.storage.GetUser(id); err != nil {
		h.logger.Println("error query user with id", id)
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the username id is not valid"})
		return
	}

	if err := h.storage.DeleteUser(id); err != nil {
		h.logger.Println("error deleting user with id", id)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "failed to delete a user"})
		return
	}

	utils.WriteJSON(rw, http.StatusCreated, types.ApiSuccess{Message: fmt.Sprintf("user deleted with id: %v", id)})
}
