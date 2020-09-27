package client

import(
	"net/http"
	"io/ioutil"
	"strings"
)

//发起http请求
func HttpPost(url string,params string) (string,error){
    client := &http.Client{}
    req, err := http.NewRequest("POST", url, strings.NewReader(params))
    if err != nil {
        return "",err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "",err
    }
    return string(body),nil
}