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

	//定时查询失败的任务重新执行todo...

	server.Run(addr)
}

