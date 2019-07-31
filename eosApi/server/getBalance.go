package server

import (
	"encoding/json"
	"eosApi/httpPost"
	"fmt"
)

func GetBalance(code string,account string,symbol string) []string{

	type getbalance struct {
		Code    string `json:"code"`
		Account string `json:"account"`
		Symbol  string `json:"symbol"`
	}

	getbalance_0:=getbalance{
		Code:code,
		Account:account,
		Symbol:symbol,
	}

	getbalance_1,_:=json.Marshal(getbalance_0)
	getbalance_2:=string(getbalance_1)

	//获取返回信息
	body :=httpPost.HttpPost(getbalance_2,"chain","get_currency_balance")
	fmt.Println("getbalance返回信息为：",string(body))

	var balance []string
	json.Unmarshal(body,&balance)

	return balance
}
