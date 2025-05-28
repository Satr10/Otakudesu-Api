package handlers

type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
