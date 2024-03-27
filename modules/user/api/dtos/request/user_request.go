package request

type UserUpdateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
