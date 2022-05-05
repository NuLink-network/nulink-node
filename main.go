package main

import (
	"github.com/NuLink-network/nulink-node/config"
	"github.com/NuLink-network/nulink-node/resource/db"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/router"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.Init()
	log.Init(viper.GetString("env"), viper.GetString("logDir"))
	dbConf := viper.GetStringMapString("database")
	db.Init(dbConf["user"], dbConf["password"], dbConf["host"], dbConf["port"], dbConf["name"])

	basicAuth := viper.GetStringMapString("basicAuth")
	accounts := gin.Accounts{
		basicAuth["user"]: basicAuth["password"],
	}

	engine := gin.Default()
	engine.Use(gin.BasicAuth(accounts))
	router.Register(engine)

	endless.ListenAndServe(viper.GetString("address"), engine)
}
