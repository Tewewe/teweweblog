package main

import (
	"teweweblog/model"
	"teweweblog/routes"
)

func main(){
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}



