package request

type GetUser struct {
	Login    string `url:"login"`
	Password string `url:"password"`
}
