package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"puhai/server/global"
	"puhai/server/initalize"
	"puhai/server/middlewares"
	"puhai/server/routers"
)

func main() {
	global.PH_DB = initalize.GormMysql()
	initalize.MysqlTables(global.PH_DB)
	r := gin.Default()
	r.Use(middlewares.Cors())
	routers.LoadRouter(r)
	err := r.Run(":8081")
	if err != nil {
		log.Fatal("run err")
	}
}