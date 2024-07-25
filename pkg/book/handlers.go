package book

import (
	"log"
	"net/http"
)

type Handler struct {
	logger  *log.Logger
	storage Storage
}

func NewHandler(logger *log.Logger, storage Storage) *Handler {
	return &Handler{logger, storage}
}

func (h *Handler) ListAll(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) ListSingle(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Create(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Update(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(rw http.ResponseWriter, r *http.Request) {

}
