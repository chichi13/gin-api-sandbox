package schemas

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
