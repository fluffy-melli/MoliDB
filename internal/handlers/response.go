package handlers

import "github.com/gin-gonic/gin"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func SendResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Status: statusCode,
		Data:   data,
	})
}

func SendErrResponse(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, ErrResponse{
		Status: statusCode,
		Error:  err,
	})
}
