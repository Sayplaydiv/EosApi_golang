package newAccountToBin

import (
	"encoding/json"
	"eosApi/config"
	"eosApi/httpPost"
	"fmt"
)





func JsonToBin(name string,owner_key string,active_key1 string,active_key2 string)string{

	var configMap = config.InitConfig("src/eosApi/config/config.ini")
	var super_account =configMap["super_account"]

	type newAccount struct {
		Code   string `json:"code"`
		Action string `json:"action"`
		Args   interface{} `json:"args"`
	}

	type Args  struct {
		Creator string `json:"creator"`
		Name    string `json:"name"`
		Owner   interface{} `json:"owner"`
		Active  interface{} `json:"active"`
	}


	type Owner  struct {
		Threshold int `json:"threshold"`
		Keys     []interface{} `json:"keys"`
		Accounts []interface{} `json:"accounts"`
		Waits    []interface{} `json:"waits"`
	}

	   type KeysOwner  struct {
		Key    string `json:"key"`
		Weight int    `json:"weight"`
	}

	type Active struct {
		Threshold int `json:"threshold"`
		Keys     []interface{} `json:"keys"`
		Accounts []interface{} `json:"accounts"`
		Waits    []interface{} `json:"waits"`
	}

	type KeysActive struct {
		Key    string `json:"key"`
		Weight int    `json:"weight"`
	}

	    KeysActive_0:=KeysActive{
		   Key:active_key1,
		   Weight:1,
	   }

	    KeysActive_1:=KeysActive{
		Key:active_key2,
		Weight:1,
	}


		KeysOwner_0:=KeysOwner{
		Key:owner_key,
		Weight:1,
	}

	Owner_0:=Owner{
		Threshold:1,
		Keys:[]interface{}{KeysOwner_0},
		Accounts:[]interface{}{},
		Waits:[]interface{}{},
	}

	Active_0:=Active{
		Threshold:2,
		Keys:[]interface{}{KeysActive_0,KeysActive_1},
		Accounts:[]interface{}{},
		Waits:[]interface{}{},
	}


	Args_0:=Args{
		Creator:super_account,
		Name:name,
		Owner:Owner_0,
		Active:Active_0,
	}

	newAccount_0 :=newAccount{
		Code:"eosio",
		Action:"newaccount",
		Args:Args_0,
	}

	newAccount_1,_:=json.Marshal(newAccount_0)
	newaccount:=string(newAccount_1)

	fmt.Println("newaccount请求数据:",newaccount)

	body:=httpPost.HttpPost(newaccount,"chain","abi_json_to_bin")


	//fmt.Println("abi_json_to_bin返回数据为：",string(body))

	type binargs struct {
		Binargs string `json:"binargs"`
	}

	var binargs_respon binargs
	json.Unmarshal(body,&binargs_respon)
	fmt.Println(string(body))
	return binargs_respon.Binargs
}

