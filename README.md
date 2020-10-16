# asyncCall
for vdongchina async api callback
请求实例
xxx.com/timer  post application/json
{
	"delay":"3600s"/"5m"/"1h"最大支持小时为单位 表示延迟执行的时间
	"request_url":"https://xxx.com/xxx" 用于执行回调的url
	"request_params":{"name":"rpc"} json格式参数列表或对象
}
1：定时回调给定的url以及参数。
2：定时检查失败的和未进行的任务，如果达到执行条件，则立即执行。