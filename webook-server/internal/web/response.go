package web

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"'msg'"`
	Data any    `json:"data"`
}
