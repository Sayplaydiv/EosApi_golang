package lock

import (
	"encoding/json"
	"eosApi/httpPost"
	"fmt"
)



func  Lock(walletName string)bool{
    lock_json:=walletName
    fmt.Println("lock加锁请求数据：",lock_json)
	body:=httpPost.HttpPost(lock_json,"wallet","lock")

	type lock struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Error   struct {
			Code    int    `json:"code"`
			Name    string `json:"name"`
			What    string `json:"what"`
			Details []struct {
				Message    string `json:"message"`
				File       string `json:"file"`
				LineNumber int    `json:"line_number"`
				Method     string `json:"method"`
			} `json:"details"`
		} `json:"error"`
	}
	var lock_0 lock
	json.Unmarshal(body,&lock_0)


	if  lock_0.Error.Code==3120002{
		fmt.Println("账户加锁失败")
		return false
	}
	fmt.Println("加锁成功")
	return true
}

func Unlock(walletName string,walletPassword string)bool{


	unlock_json_0:=[]interface{}{walletName,walletPassword}
	unlock_json_1,_:=json.Marshal(unlock_json_0)
	unlock_json :=string(unlock_json_1)
	fmt.Println("unlock解锁请求数据：",unlock_json)
	body:=httpPost.HttpPost(unlock_json,"wallet","unlock")
	//fmt.Println(string(body))

	type unlock struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Error   struct {
			Code    int    `json:"code"`
			Name    string `json:"name"`
			What    string `json:"what"`
			Details []struct {
				Message    string `json:"message"`
				File       string `json:"file"`
				LineNumber int    `json:"line_number"`
				Method     string `json:"method"`
			} `json:"details"`
		} `json:"error"`
	}
	var unlock_0 unlock
	json.Unmarshal(body,&unlock_0)
	if unlock_0.Error.Code==3120007 {
		fmt.Println("账户已解锁")
		return true
	}else if unlock_0.Error.Code==0 {
		fmt.Println("部署账户解锁成功")
		return true
	}
	return false

}
