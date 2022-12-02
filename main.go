package main

import (
	"education/config"
	"education/database"
	"education/middle"
	"education/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "conf/config_release.yaml", "input the config")
)

func init() {
	pflag.Parse()
	if err := config.InitConfig(*cfg); err != nil {
		panic(err)
	}
	database.InitMysql()
	database.InitRedis()
}

func main() {
	g := gin.Default()
	g.Use(middle.Cors())
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	router.InitRouter(g)
	g.Run(viper.GetString("addr"))
}
