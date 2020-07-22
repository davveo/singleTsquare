package request

type UserRequest struct {
	UserName       string `json:"username" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeatpassword" binding:"required"`
	Code           string `json:"code" binding:"required"`
}
