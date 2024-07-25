package types

type ApiError struct {
	Error string `json:"error"`
}
type ApiSuccess struct {
	Message string `json:"message"`
}
