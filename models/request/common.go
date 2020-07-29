package request

type LoginRequestJson struct {
	LoginId string `json:"login_id"  binding:"required"`
	Mode    string `json:"mode"`
}
