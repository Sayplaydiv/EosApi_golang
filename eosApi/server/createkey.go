package server

import (
	"encoding/json"
	"eosApi/httpPost"
	"eosApi/server/lock"
	"fmt"
)

func CreateKeys(walletName string,walletPassword string) (string){
	isTrue:=lock.Unlock(walletName,walletPassword)

	if  isTrue {
		wallet_name:=walletName
		encry_mode:="k1"

		createKey_json:=[]string{wallet_name,encry_mode}
		createKey_json_0,_:=json.Marshal(createKey_json)
		createKey_json_1:=string(createKey_json_0)

		body:=httpPost.HttpPost(createKey_json_1,"wallet","create_key")

		fmt.Println("新建key的公钥为：",string(body))


		isTrue:=lock.Lock(walletName)
		if isTrue==false {
			fmt.Println("加锁失败")
		}

		return string(body)

	}else {
		fmt.Println("解锁失败")
	}

        return ""
}
