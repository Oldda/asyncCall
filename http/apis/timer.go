package apis

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"asyncCall/http/client"
	"time"
    "asyncCall/db"
)

//延时消费回调
func HandleTimer(w http.ResponseWriter, r *http.Request){
	//获取参数
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
    	w.Write([]byte("接收参数出错了"))
        return
    }

    //json to map
    var params map[string]interface{}
    if err = json.Unmarshal(body, &params); err != nil {
        w.Write([]byte("参数格式有错误"))
        return
    }
   	log.Println(params)
    
    timeStr,_ := params["delay"].(string)
    requestUrl,_ := params["request_url"].(string)
    requestParams,_ := json.Marshal(params["request_params"])

    //存储任务
    job := db.Jobs{Delay:timeStr,RequestUrl:requestUrl,RequestParams:string(requestParams),CreateTime:uint(time.Now().Unix())}
    db.MysqlEngine.Create(&job)

    //请求
    go exec(timeStr,requestUrl,string(requestParams),job.ID)
	w.Write([]byte("任务已加载..."))
}

//执行任务
func exec(timeStr, requestUrl, requestParams string,jobId uint){
	duration,err := time.ParseDuration(timeStr)
	if err != nil{
		log.Println(err)
	}
	timer := time.NewTimer(duration)
	log.Println(<-timer.C,"job running")
	result,err := client.HttpPost(requestUrl,requestParams)
    if err != nil{
    	log.Println(err)
    }
    var data map[string]string
    json.Unmarshal([]byte(result),&data)
    log.Println("result_data:",data)
    job := db.Jobs{ID:jobId}
    db.MysqlEngine.First(&job)
    job.RequestTime = uint(time.Now().Unix())
    if data["code"] == "200" || data["code"] == "202"{
        job.Status = 1
    }else{
        job.Status = 2
    }
    db.MysqlEngine.Save(&job)
    log.Println("job done")
}
