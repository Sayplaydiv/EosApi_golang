package delegatebwToBin

import (
	"encoding/json"
	"eosApi/config"
	"eosApi/httpPost"
	"fmt"
)

func JsonToBin(receiver string,StakeNetQuantity string,stake_cpu_quantity string)string {

	var configMap = config.InitConfig("src/eosApi/config/config.ini")
	var super_account =configMap["super_account"]

	type delegatebw struct {
		Code   string `json:"code"`
		Action string `json:"action"`
		Args   interface{} `json:"args"`
	}

	type Args  struct {
		From             string `json:"from"`
		Receiver         string `json:"receiver"`
		StakeNetQuantity string `json:"stake_net_quantity"`
		StakeCPUQuantity string `json:"stake_cpu_quantity"`
		Transfer         int    `json:"transfer"`
	}

	   Args_0 := Args{
		   From:super_account,
		   Receiver:receiver,
		   StakeNetQuantity:StakeNetQuantity,
		   StakeCPUQuantity:stake_cpu_quantity,
		   Transfer:0,

	}

	delegatebw_0:=delegatebw{
		Code:"eosio",
		Action:"delegatebw",
		Args:Args_0,
	}

	delegatebw_1,_:=json.Marshal(delegatebw_0)

	delegatebw_2:=string(delegatebw_1)
	fmt.Println("抵押cpu和net请求数据：",delegatebw_2)

	body:=httpPost.HttpPost(delegatebw_2,"chain","abi_json_to_bin")

	type binargs struct {
		Binargs string `json:"binargs"`
	}

	var binargs_respon binargs
	json.Unmarshal(body,&binargs_respon)
	fmt.Println(string(body))
	return binargs_respon.Binargs

}
