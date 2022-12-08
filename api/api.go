package api

type (
	Response struct {
		StatusCode int         `json:"code"`
		Msg        string      `json:"msg"`
		Data       interface{} `json:"data"`
	}

	List struct {
		Page  int `json:"page"`
		Size  int `json:"size"`
		Total int `json:"total"`
	}
)

func Success(data interface{}, param ...interface{}) Response {

	return Response{}
}

func Fail(code int, msg ...string) Response {

	return Response{}
}
