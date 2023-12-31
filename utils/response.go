package utils

type HttpResponse struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetResponse(status uint, message string, data interface{}) HttpResponse {
	res := HttpResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return res
}
