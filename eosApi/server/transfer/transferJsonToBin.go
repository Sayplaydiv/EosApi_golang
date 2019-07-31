package transfer

import (
	"encoding/json"
	"eosApi/httpPost"
	"fmt"
)

func JsonToBin(from_address string,to_address string,quantity string,memo string)string{

	type transfer struct {
		Code   string `json:"code"`
		Action string `json:"action"`
		Args   struct {
			From     string `json:"from"`
			To       string `json:"to"`
			Quantity string `json:"quantity"`
			Memo     string `json:"memo"`
		} `json:"args"`
	}

	type args_0  struct{
		From     string `json:"from"`
		To       string `json:"to"`
		Quantity string `json:"quantity"`
		Memo     string `json:"memo"`
	}


	transfer_0:=transfer{
		Code:"eosio.token",
		Action:"transfer",
		Args:args_0{From:from_address,To:to_address,Quantity:quantity,Memo:memo},
	}

	transfer_1,_:=json.Marshal(transfer_0)
	transfer_2:=string(transfer_1)

	fmt.Println("transfer请求数据：",transfer_2)
	body:=httpPost.HttpPost(transfer_2,"chain","abi_json_to_bin")


	//fmt.Println("abi_json_to_bin返回数据为：",string(body))

	type binargs struct {
		Binargs string `json:"binargs"`
	}

	var binargs_respon binargs
	json.Unmarshal(body,&binargs_respon)

	return binargs_respon.Binargs

}
