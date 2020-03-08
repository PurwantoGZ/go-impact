package request

//LoginRequest for wrap login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
