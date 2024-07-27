// Package classification of [BOOKNEST] API
//
// Documentation for [BOOKNEST] API
//
//	Schemes: http
//	BasePath: /api
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package docs

import (
	"github.com/shahinrahimi/booknest/pkg/book"
	"github.com/shahinrahimi/booknest/types"
)

// swagger:response errorResponse
type ApiErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body types.ApiError
}

// swagger:response successResponse
type ApiSuccessResponseWrapper struct {
	// Description of the success
	// in: body
	Body types.ApiSuccess
}

// A list of books
// swagger:response booksResponse
type BooksResponseWrapper struct {
	// All current books
	// in: body
	Body []book.Book
}

// a single book
// swagger:response bookResponse
type BookResponseWrapper struct {
	// Newly created book
	// in: body
	Body book.Book
}

// swagger:parameters updateBook createBook
type BookParamsWrapper struct {
	// Book data structure to Update or Create.
	// Note: the id filed is ignored by update and create operations
	// in: body
	// required: true
	Body book.Book
}

// swagger:parameters updateBook deleteBook listSingleBook
type BookIDParamsWrapper struct {
	// The id of the book for which the operation related
	// in: path
	// required: true
	ID string `json:"id"`
}
