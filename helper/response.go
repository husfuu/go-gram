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

// this is for swagger needs
type ExampleErrorResponse struct {
	Message  string `json:"message" example:"user bad request"`
	Email    string `json:"email" example:"cannot be empty"`
	Password string `json:"password" example:"cannot be empty"`
}
