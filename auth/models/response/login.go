package response

type LoginResponse struct {
	Username string `json:"username"`
	UserId   string `json:"userId"`
	Token
}
