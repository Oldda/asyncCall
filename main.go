package main

import(
	"asyncCall/http/server"
	"asyncCall/http/utils"
	"asyncCall/http/apis"
	_ "asyncCall/db"
)

func main(){

	addr := utils.GetAddrFromFlag()

	server := server.NewServerMux()

	server.POST("/timer",apis.HandleTimer) //注册定时消费路由

	go server.Ticker("3s",apis.ReExecJob) //定义每分钟需要执行一次处理函数

	server.Run(addr)
}

