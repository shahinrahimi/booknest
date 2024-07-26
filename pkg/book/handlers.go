package book

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
	books, err := h.storage.GetBooks()
	if err != nil {
		h.logger.Println("error query user from DB", err)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}
	if len(books) == 0 {
		h.logger.Println("no book found in DB")
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiSuccess{Message: "there is no book found"})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, books)
}

func (h *Handler) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	b, err := h.storage.GetBook(id)
	if err != nil {
		h.logger.Println("error query book with id", id)
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the book id is not valid"})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, b)
}

func (h *Handler) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch validated book from the context
	b := r.Context().Value(KeyBook{}).(Book)
	// create a new book
	newBook := NewBook(b.Title, b.Author, b.Description, b.Cover, b.Price)

	if err := h.storage.CreateBook(newBook); err != nil {
		h.logger.Println("error creating a new book")
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "failed to create a new book"})
		return
	}
	utils.WriteJSON(rw, http.StatusCreated, types.ApiSuccess{Message: fmt.Sprintf("new book created with id: %v", newBook.ID)})

}

func (h *Handler) Update(rw http.ResponseWriter, r *http.Request) {
	// fetch validated book from the context
	b := r.Context().Value(KeyBook{}).(Book)
	// get book id
	id := mux.Vars(r)["id"]
	// check if book exists on the DB
	fetchedBook, err := h.storage.GetBook(id)
	if err != nil {
		h.logger.Println("error query book with id", id)
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the book id is not valid"})
		return
	}
	// assign ID to validated book
	b.ID = fetchedBook.ID

	h.storage.UpdateBook(id, &b)

}

func (h *Handler) Delete(rw http.ResponseWriter, r *http.Request) {

}
