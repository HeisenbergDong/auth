package main

import (
	"auth/gin"
	"auth/global"
	"auth/grpc/auth/start"
	"auth/pkg/gorm"
	"auth/pkg/redis"
	"auth/pkg/viper"
	"auth/pkg/zap"
	"auth/router"
)

func main() {
	global.VIPER = viper.InitViper()     // 初始化Viper配置文件
	global.LOG = zap.InitZap()           // 初始化zap日志库
	global.DB = gorm.InitGorm()          // 初始化gorm连接数据库
	global.REDIS = redis.InitRedis()     // 初始化redis服务
	global.ROUTER = router.InitRouters() // 初始化路由
	gin.Run()                            // 启动GIN服务
	grpcServer := start.NewGrpcServer()
	grpcServer.Run()                     // 启动GRPC服务
	select {}
}
