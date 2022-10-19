package helper

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

func NewResponse(status int, data interface{}, error error) *Response {

	if error != nil {
		return &Response{
			Status: status,
			Data:   data,
			Error:  map[string]interface{}{"message": error.Error()},
		}
	}

	return &Response{
		status,
		data,
		error,
	}
}
