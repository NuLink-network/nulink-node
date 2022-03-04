package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/router"
)

func main() {
	engine := gin.Default()

	router.Register(engine)

	log.Panic(engine.Run())
}
