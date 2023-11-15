package response

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
}

type JsonGetResponse struct {
	Count   uint64 `json:"count"`
	Status  uint16 `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GetJsonResponse(c *gin.Context, count uint32, statusCode uint16, message string, data any) {
	res := JsonGetResponse{
		Count:   uint64(count),
		Status:  statusCode,
		Message: message,
		Data:    data,
	}

	c.JSON(int(statusCode), res)
}

func GenericJsonResponse(c *gin.Context, statusCode uint16, message string) {
	res := map[string]any{
		"status":  statusCode,
		"message": message,
	}

	c.JSON(int(statusCode), res)
}
