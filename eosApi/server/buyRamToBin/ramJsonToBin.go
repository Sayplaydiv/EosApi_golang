package buyRamToBin

import (
	"encoding/json"
	"eosApi/config"
	"eosApi/httpPost"
	"fmt"
)

func JsonToBin(receiver string,quant string)string{

	var configMap = config.InitConfig("src/eosApi/config/config.ini")
	var super_account =configMap["super_account"]

	type buyram_0 struct {
		Code   string `json:"code"`
		Action string `json:"action"`
		Args   struct{
			Payer    string `json:"payer"`
			Receiver string `json:"receiver"`
			Quant    string `json:"quant"`
		} `json:"args"`
	}

	type args_0  struct{
		Payer    string `json:"payer"`
		Receiver string `json:"receiver"`
		Quant    string `json:"quant"`
	}

	buyram_1 :=buyram_0{
		"eosio",
		"buyram",
		args_0{Payer:super_account,Receiver:receiver,Quant:quant,},
	}

	buyram_2,_:=json.Marshal(buyram_1)
	buyram :=string(buyram_2)


	fmt.Println("buyram请求数据：",buyram)
	body:=httpPost.HttpPost(buyram,"chain","abi_json_to_bin")


	//fmt.Println("abi_json_to_bin返回数据为：",string(body))

	type binargs struct {
		Binargs string `json:"binargs"`
	}

	var binargs_respon binargs
	json.Unmarshal(body,&binargs_respon)

	return binargs_respon.Binargs
}

