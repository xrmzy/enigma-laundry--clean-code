package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

type MultipleResponse struct {
	Status Status        `json:"status"`
	Data   []interface{} `json:"data"`
}

func SendSingleResponse(c *gin.Context, code int, description string, data interface{}) {
	c.JSON(http.StatusCreated, SingleResponse{
		Status: Status{
			Code:        code,
			Description: description,
		},
		Data: data,
	})
}
