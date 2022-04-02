package resp

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
		Code: CodeSuccess,
		Msg:  MsgSuccess,
		Data: EmptyStruct,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  MsgSuccess,
		Data: data,
	})
}

func Error(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  code2msg[code],
		Data: EmptyStruct,
	})
}

func Resp(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  code2msg[code],
		Data: data,
	})
}

func ParameterErr(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeParameterErr,
		Msg:  code2msg[CodeParameterErr],
		Data: EmptyStruct,
	})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: CodeInternalServerError,
		Msg:  MsgInternalServerError,
		Data: EmptyStruct,
	})
}
