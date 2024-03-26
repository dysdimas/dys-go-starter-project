package response

type AuthResponse struct {
	Token     string `json:"token"`
	CreatedAt int64  `json:"created_at"`
	ExpiredAt int64  `json:"expired_at"`
}
