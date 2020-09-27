# asyncCall
for vdongchina async api callback
请求实例
xxx.com/timer  post application/json
{
	"delay":"3600s"/"5m"/"1h"最大支持小时为单位 表示延迟执行的时间
	"request_url":"https://xxx.com/xxx" 用于执行回调的url
	"request_params":{"name":"rpc"} json格式参数列表或对象
}