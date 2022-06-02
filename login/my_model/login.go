package my_model
type LoginAck struct {
	Code int `json:"code"`
	Token string `json:"token"`
}