package request

type UserRequest struct {
	UserName       string `json:"username" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeat_password" binding:"required"`
	Code           string `json:"code" binding:"required"`
}

type LoginRequest struct {
	LoginId  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type FastLoginRequest struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

type BindRequest struct {
	Phone      string `json:"phone"`
	Code       string `json:"code"`
	LoginId    string `json:"login_id"`
	Password   string `json:"password"`
	IdentifyId string `json:"identify_id" binding:"required"`
}

type VerifyCodeRequest struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
