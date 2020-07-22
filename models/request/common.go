package request

type PhoneRequestJson struct {
	Phone string `json:"phone"  binding:"required"`
}
