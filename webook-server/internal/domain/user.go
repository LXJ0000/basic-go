package domain

type User struct {
	UserId   int64  `json:"user_id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Password string `json:"-"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Intro    string `json:"intro"`
	WebSite  string `json:"web_site"`
}
