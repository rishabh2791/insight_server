package views

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}
