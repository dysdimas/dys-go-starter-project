package response

type RegisterResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
