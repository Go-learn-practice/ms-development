package common

type BusinessCode int
type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}

// Success 响应成功
func (r *Result) Success(data interface{}) *Result {
	r.Code = 200
	r.Msg = "success"
	r.Data = data
	return r
}

// Fail 响应失败
func (r *Result) Fail(code BusinessCode, msg string) *Result {
	r.Code = code
	r.Msg = msg
	return r
}
