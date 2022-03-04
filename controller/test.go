package controller

import (
	"github.com/NuLink-network/nulink-node/utils"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	utils.SuccessNil(c)
}
