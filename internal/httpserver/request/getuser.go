package request

type GetUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
