package response

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   *Error `json:"error,omitempty"`
}

func (r *BaseResponse) NewError(code, message string) {
	r.Error = &Error{
		Code:    code,
		Message: message,
	}
}

func (r *BaseResponse) SetSuccessStatus(isSuccess bool) {
	r.Success = isSuccess
}
