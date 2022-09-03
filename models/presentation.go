package models

type ResponseErr struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

type ResponseSuccess struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

// PresentErr is a method to present error message
func (r *ResponseErr) PresentErr() *ResponseErr {
	return &ResponseErr{
		Status: false,
		Error:  r.Error,
	}
}

// PresentSuccess is method to present success response
func (r *ResponseSuccess) PresentSuccess() *ResponseSuccess {
	return &ResponseSuccess{
		Status: true,
		Error:  "",
	}
}
