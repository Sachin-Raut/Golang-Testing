package errors

//APIError is
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
