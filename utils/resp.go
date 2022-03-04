package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var EmptyStruct = struct{}{}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessNil(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: SuccessCode,
		Msg:  SuccessMsg,
		Data: EmptyStruct,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: SuccessCode,
		Msg:  SuccessMsg,
		Data: data,
	})
}
