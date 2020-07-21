package request

type UserRequest struct {
	NickName string `json:"nickname"  `
	UserName string `json:"username"  binding:"required"`
	Pssword  string `json:"password" binding:"required"`
}
