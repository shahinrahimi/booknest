package book

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

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

// swagger:route GET /book book listBooks
// Return a list of books from database
// responses:
//	200: booksResponse
//	500: errorResponse

// ListAll handles GET requests
func (h *Handler) ListAll(rw http.ResponseWriter, r *http.Request) {
	books, err := h.storage.GetBooks()
	if err != nil {
		h.logger.Println("error query user from DB", err)
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}
	if len(books) == 0 {
		h.logger.Println("no book found in DB")
		utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: "there is no book found"})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, books)
}

// swagger:route GET /book/{id} book listSingleBook
// Return a book from database
// responses:
//	200: bookResponse
//	404: errorResponse
//  500: errorResponse

// ListSingle handles GET requests
func (h *Handler) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	b, err := h.storage.GetBook(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: "the book id is not valid"})
			return
		} else {
			h.logger.Println("error query book with id", id)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the book id is not valid"})
			return
		}

	}
	utils.WriteJSON(rw, http.StatusOK, b)
}

// swagger:route POST /api/book/ book createBook
// Create a new Book
// responses:
//	201: successResponse
//	400: errorResponse
//  500: errorResponse

// Create handles POST request
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

// swagger:route PUT /api/book/{id} book updateBook
// update the Book
// responses:
//	201: successResponse
//	400: errorResponse
//	404: errorResponse
//  500: errorResponse

// Update handle PUT request
func (h *Handler) Update(rw http.ResponseWriter, r *http.Request) {
	// fetch validated book from the context
	b := r.Context().Value(KeyBook{}).(Book)
	// get book id
	id := mux.Vars(r)["id"]
	// check if book exists on the DB
	fetchedBook, err := h.storage.GetBook(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: "the book id is not valid"})
			return
		} else {
			h.logger.Println("error query book with id", id)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the book id is not valid"})
			return
		}
	}
	// internal data
	b.UpdatedAt = time.Now().UTC()
	b.CreatedAt = fetchedBook.CreatedAt

	if err := h.storage.UpdateBook(id, &b); err != nil {
		h.logger.Println("error update the book with id", id)
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the book id is not valid"})
		return
	}

	utils.WriteJSON(rw, http.StatusBadRequest, types.ApiSuccess{Message: fmt.Sprintf("the book updated with id: %s", id)})

}

// swagger:route DELETE /api/book/{id} book deleteBook
// delete the Book
// responses:
//	201: successResponse
//	400: errorResponse
//	404: errorResponse
//  500: errorResponse

// Delete handle DELETE request
func (h *Handler) Delete(rw http.ResponseWriter, r *http.Request) {
	// get book id
	id := mux.Vars(r)["id"]
	// check if book exist on the db
	_, err := h.storage.GetBook(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: "the book id is not valid"})
			return
		} else {
			h.logger.Println("error query book with id", id)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the book id is not valid"})
			return
		}
	}
	if err := h.storage.DeleteBook(id); err != nil {
		h.logger.Println("error delete the book with id", id)
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "the book id is not valid"})
		return
	}

	utils.WriteJSON(rw, http.StatusCreated, types.ApiSuccess{Message: fmt.Sprintf("the book deleted with id: %s", id)})

}
