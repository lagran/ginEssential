package main

import (
	"funboy.top/ginessential/common"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
