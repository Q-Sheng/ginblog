package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	// 连接数据库
	model.InitDb()

	routes.InitRouter()

}
