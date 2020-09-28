package apis

import(
	"asyncCall/db"
	"time"
)

func ReExecJob(){
	var jobs []db.Jobs
	//查询crm_jobs表中状态未2（执行失败）和（状态为0 && create_time + delay <= now
	db.MysqlEngine.Where("status = ?","2").Or("status = ? AND create_time + delay <= ?","0", time.Now().Unix()).Find(&jobs)
	for _,job := range jobs{
		//分别处理每一条
		duration,_ := time.ParseDuration(job.Delay)
		go exec(duration, job.RequestUrl, job.RequestParams,job.ID)
	}
}