package main

import (
	"fmt"
	"github.com/winterssy/sreq"
	"time"
)

var API_URL string = "http://127.0.0.1:5700"
var TAG_API_URL string = "http://127.0.0.1:5001/get_info"

type Data struct {
	Total    float32 `json:"total"`
	Used     float32 `json:"used"`
	Left     float32 `json:"left"`
	Download float32 `json:"download"`
	Update   float32 `json:"update"`
}

func GetGroupList() map[string]interface{} {
	res := sreq.Get(API_URL + "/get_group_list")
	fmt.Println(res)
	data := make(map[string]interface{})
	_ = res.JSON(&data)
	return data
}

func SendMsgToGroup(groupId int64, msg string, raw bool) {
	params := sreq.Params{
		"group_id":    groupId,
		"message":     msg,
		"auto_escape": raw,
	}
	_ = sreq.Get(API_URL+"/send_group_msg", sreq.WithQuery(params))
}

func GetTagData(data *Data) {
	err := sreq.Get(TAG_API_URL).JSON(&data)
	if err != nil {
		return
	}
}

func main_qqbot() {
	oldData := Data{}
	ticker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-ticker.C:
			newData := Data{}
			GetTagData(&newData)
			fmt.Printf("%.2f\n", newData.Used-oldData.Used)
			fmt.Printf("%.2f, %.2f\n", newData.Used, oldData.Used)
			if oldData.Total != 0 && newData.Used-oldData.Used > 0.5 {
				println("以发送警告")
				msg := fmt.Sprintf("流量异常提醒\n订阅流量: %.2fG\n已用流量: %.2fG\n剩余流量: %.2f%s",
					newData.Total, newData.Used, newData.Left, "%",
				)
				SendMsgToGroup(263272274, msg, false)
			}
			oldData.Total = newData.Total
			oldData.Used = newData.Used
			oldData.Left = newData.Left
			oldData.Download = newData.Download
			oldData.Update = newData.Update
		}
	}

}
