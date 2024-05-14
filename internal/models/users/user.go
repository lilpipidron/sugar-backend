package users

type User struct {
	UserID   int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	UserInfo UserInfo
}
