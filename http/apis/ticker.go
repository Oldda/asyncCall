package apis

import(
	"log"
)

func ReExecJob(){
	//查询crm_jobs表中状态未2（执行失败）和（状态为0 && create_time + delay <= now）
	//分别根据该任务的url和params再次执行一次。
	//code by zl
	log.Println("test here")
}