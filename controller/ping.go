package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/resp"
)

func Ping(c *gin.Context) {
	resp.SuccessNil(c)
}
