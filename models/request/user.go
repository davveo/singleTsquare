package request

type UserRequest struct {
	UserName string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Code     string `json:"code" binding:"required"`
}
