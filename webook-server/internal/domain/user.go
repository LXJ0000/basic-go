package domain

type User struct {
	UserId   int64  `json:"user_id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"-"`
}
