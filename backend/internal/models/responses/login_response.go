package response

import "github.com/goccy/go-json"

type LoginResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	*BaseResponse
}

func NewLoginResponse() *LoginResponse {
	return &LoginResponse{BaseResponse: &BaseResponse{}}
}

func NewLoginResponseWithError(code, message string) ([]byte, error) {
	return NewLoginResponse().SetError(code, message).ToByte()
}

func NewLoginResponseWithAccessToken(token string) ([]byte, error) {
	return NewLoginResponse().SetAccessToken(token).SetSuccessStatus(true).ToByte()
}

func (r *LoginResponse) SetError(code, message string) *LoginResponse {
	r.BaseResponse.NewError(code, message)
	return r
}

func (r *LoginResponse) SetAccessToken(token string) *LoginResponse {
	r.AccessToken = token
	return r
}

func (r *LoginResponse) SetSuccessStatus(isSuccess bool) *LoginResponse {
	r.BaseResponse.SetSuccessStatus(isSuccess)
	return r
}

func (r *LoginResponse) ToByte() ([]byte, error) {
	return json.Marshal(r)
}
