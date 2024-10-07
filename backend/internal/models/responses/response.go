package response

type Status string

const (
	Error   Status = "error"
	Success Status = "success"
)

type BaseResponse[T any] struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    T      `json:"data,omitempty"`
}
