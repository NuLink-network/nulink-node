package router

import (
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/controller"
)

const (
	TestGroup = "/test"
)

func Register(engine *gin.Engine) {
	test := engine.Group(TestGroup)
	test.GET("/ping", controller.Ping)
}
