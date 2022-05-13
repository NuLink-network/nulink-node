package main

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/NuLink-network/nulink-node/config"
	"github.com/NuLink-network/nulink-node/resource/db"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/router"
)

func main() {
	config.Init()
	log.Init(viper.GetString("env"), viper.GetString("logDir"))
	dbConf := viper.GetStringMapString("database")
	db.Init(dbConf["user"], dbConf["password"], dbConf["host"], dbConf["port"], dbConf["name"])

	//basicAuth := viper.GetStringMapString("basicAuth")
	//accounts := gin.Accounts{
	//	basicAuth["user"]: basicAuth["password"],
	//}

	engine := gin.Default()
	//engine.Use(gin.BasicAuth(accounts))
	engine.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           2 * time.Minute,
	}))
	router.Register(engine)

	endless.ListenAndServe(viper.GetString("address"), engine)
}
