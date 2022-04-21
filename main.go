package main

import (
	"github.com/NuLink-network/nulink-node/config"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/router"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	//f, _ := os.Create("node-request.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	config.Init()
	log.InitLogger(logrus.DebugLevel, "./node.log")
	//db.InitDB("", "", "", "", "")

	engine := gin.Default()

	engine.Use(gin.Recovery())
	router.Register(engine)

	endless.ListenAndServe(":8088", engine)
}
