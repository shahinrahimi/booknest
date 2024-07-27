package types

type ApiError struct {
	Error string `json:"error"`
}
type ApiSuccess struct {
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
