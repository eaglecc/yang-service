package main

import (
	"fmt"
	"yang-service/config"
	"yang-service/dao"
	"yang-service/routers"
)

func main() {
	// 1. 加载数据库配置
	config.InitConfig()
	// 2. 启动MySQL数据库
	dao.ConnectMySQLDB()
	// 3. 启动路由
	r := routers.InitRouter()
	r.Run(fmt.Sprintf(":%d", config.Conf.ServerPort))
}
