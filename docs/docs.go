// Package classification of [BOOKNEST] API
//
// Documentation for [BOOKNEST] API
//
//	Schemes: http
//	BasePath: /
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
