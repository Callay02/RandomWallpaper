package utils

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Result) OK(message string) Result {
	r.Code = 200
	r.Message = message
	r.Data = ""
	return *r
}

func (r *Result) Error(message string) Result {
	r.Code = 400
	r.Message = message
	r.Data = ""
	return *r
}

func (r *Result) OKSetData(data interface{}) Result {
	r.Code = 200
	r.Message = "success"
	r.Data = data
	return *r
}

func (r *Result) OKSetMsgData(message string, data interface{}) Result {
	r.Code = 200
	r.Message = message
	r.Data = data
	return *r
}
