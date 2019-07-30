package server

import (
	"encoding/json"
	"eosApi/httpPost"
	"fmt"
)

func CreateWallet(wallet_name string) (string,string){

	fmt.Println("创建钱包请求数据为：",wallet_name)
	body:=httpPost.HttpPost(wallet_name,"wallet","create")

	type createWallet struct {
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
	var createWallet_0 createWallet
	json.Unmarshal(body,&createWallet_0)
	if createWallet_0.Code==0 {
		fmt.Println("钱包密码为：",string(body))
		fmt.Println("钱包名称为：",wallet_name)
		return wallet_name,string(body)
	}else {
		fmt.Println(string(body))
	}
       return "",""
}
