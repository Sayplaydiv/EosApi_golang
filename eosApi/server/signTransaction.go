package server

import (
	"encoding/json"
	"eosApi/config"
	"eosApi/httpPost"
	"eosApi/server/buyRamToBin"
	"eosApi/server/delegatebwToBin"
	"eosApi/server/lock"
	"eosApi/server/newAccountToBin"
	"fmt"
	"strings"
	"time"
)





func SignTransaction(method string,account_name string,owner_key string,active_key1 string,active_key2 string,buyRamAmount string,buyCpuAmount string,buyNetAmount string)([]string,[]interface{},string,int,int){


	var configMap = config.InitConfig("src/eosApi/config/config.ini")
	var chain_id =configMap["chain_id"]
    var public_key =configMap["account_public_key"]
	var super_account =configMap["super_account"]


	type sign struct {
		RefBlockNum    int    `json:"ref_block_num"`
		RefBlockPrefix int    `json:"ref_block_prefix"`
		Expiration     string `json:"expiration"`
		Actions        []interface{} `json:"actions"`
	}

	  type Actions   struct {
		Account       string `json:"account"`
		Name          string `json:"name"`
		Authorization []interface{} `json:"authorization"`
		Data string `json:"data"`
	}

	type Authorization struct {
		Actor      string `json:"actor"`
		Permission string `json:"permission"`
	}
	Authorization_0:=Authorization{
		Actor:super_account,
		Permission:"owner",
	}




	//RefBlockNum:=GetInfo()
	Expiration,RefBlockPrefix,RefBlockNum:=GetBlock()

	datetime:=strings.Replace(Expiration,"T"," ",-1)

	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05.000"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp_0 := tmp.Unix()    //转化为时间戳 类型是int64

	//时间戳加2分钟，后转回
	timestamp_1 :=timestamp_0+1200
	timestamp_2:= time.Unix(timestamp_1, 0).Format(timeLayout)
	timestamp:=strings.Replace(timestamp_2," ","T",-1)
	//fmt.Println(timestamp)


	if method=="newaccount" {
		var wallet_name =configMap["wallet_name"]
		var wallet_password =configMap["wallet_password"]
		//Authorization_1,_:=json.Marshal(Authorization_0)

		data_newAccount:=newAccountToBin.JsonToBin(account_name,owner_key,active_key1,active_key2)
		data_buyRam:=buyRamToBin.JsonToBin(account_name,buyRamAmount)
		data_delegatebw:=delegatebwToBin.JsonToBin(account_name,buyCpuAmount,buyNetAmount)

		Actions_buyram:=Actions{
			Account:"eosio",
			Name:"buyram",
			Authorization:[]interface{}{Authorization_0},
			Data:data_buyRam,
		}

		Actions_newaccount:=Actions{
			Account:"eosio",
			Name:"newaccount",
			Authorization:[]interface{}{Authorization_0},
			Data:data_newAccount,
		}

		Actions_delegatebw:=Actions{
			Account:"eosio",
			Name:"delegatebw",
			Authorization:[]interface{}{Authorization_0},
			Data:data_delegatebw,
		}

		//Actions_1,_:=json.Marshal(Actions_0)

		sign_0:=sign{
			RefBlockNum:RefBlockNum,
			RefBlockPrefix:RefBlockPrefix,
			Expiration:timestamp,
			Actions:[]interface{}{Actions_newaccount,Actions_buyram,Actions_delegatebw},
		}

		Actions_push:=[]interface{}{Actions_newaccount,Actions_buyram,Actions_delegatebw}





		sign_transaction_1:=[]interface{}{sign_0,[]string{public_key},chain_id}

		sign_transaction_2,_:=json.Marshal(sign_transaction_1)

		sign_transaction :=string(sign_transaction_2)


		// fmt.Println(reflect.TypeOf(sign_0))


		fmt.Println("sign_transaction请求信息：",sign_transaction)


		isTrue:=lock.Unlock(wallet_name,wallet_password)
		if isTrue {
			body :=httpPost.HttpPost(sign_transaction,"wallet","sign_transaction")
			fmt.Println("sign_transaction签名返回数据：",string(body))


			type sign_respon struct {
				Expiration         string        `json:"expiration"`
				RefBlockNum        int           `json:"ref_block_num"`
				RefBlockPrefix     int           `json:"ref_block_prefix"`
				MaxNetUsageWords   int           `json:"max_net_usage_words"`
				MaxCPUUsageMs      int           `json:"max_cpu_usage_ms"`
				DelaySec           int           `json:"delay_sec"`
				ContextFreeActions []interface{} `json:"context_free_actions"`
				Actions            []struct {
					Account       string `json:"account"`
					Name          string `json:"name"`
					Authorization []struct {
						Actor      string `json:"actor"`
						Permission string `json:"permission"`
					} `json:"authorization"`
					Data string `json:"data"`
				} `json:"actions"`
				TransactionExtensions []interface{} `json:"transaction_extensions"`
				Signatures            []string      `json:"signatures"`
				ContextFreeData       []interface{} `json:"context_free_data"`
			}
			var sign_respon_0 sign_respon
			json.Unmarshal(body,&sign_respon_0)
			fmt.Println("获取签名Signatures：",sign_respon_0.Signatures)


			return sign_respon_0.Signatures,Actions_push,timestamp,RefBlockNum,RefBlockPrefix
		}else {
			fmt.Println("解锁失败")
		}

	}else if method=="buyram"{
		var wallet_name =configMap["wallet_name"]
		var wallet_password =configMap["wallet_password"]
		data_buyRam:=buyRamToBin.JsonToBin(account_name,buyRamAmount)

		fmt.Println("获取编码后的data数据:",data_buyRam)
		//Authorization_1,_:=json.Marshal(Authorization_0)


		Actions_buyram:=Actions{
			Account:"eosio",
			Name:"buyram",
			Authorization:[]interface{}{Authorization_0},
			Data:data_buyRam,
		}


		//Actions_1,_:=json.Marshal(Actions_0)

		sign_0:=sign{
			RefBlockNum:RefBlockNum,
			RefBlockPrefix:RefBlockPrefix,
			Expiration:timestamp,
			Actions:[]interface{}{Actions_buyram},
		}

		Actions_push:=[]interface{}{Actions_buyram}



		sign_transaction_1:=[]interface{}{sign_0,[]string{public_key},chain_id}

		sign_transaction_2,_:=json.Marshal(sign_transaction_1)

		sign_transaction :=string(sign_transaction_2)


		// fmt.Println(reflect.TypeOf(sign_0))


		isTrue:=lock.Unlock(wallet_name,wallet_password)
		if isTrue {
			body :=httpPost.HttpPost(sign_transaction,"wallet","sign_transaction")
			fmt.Println("sign_transaction签名返回数据：",string(body))


			type sign_respon struct {
				Expiration         string        `json:"expiration"`
				RefBlockNum        int           `json:"ref_block_num"`
				RefBlockPrefix     int           `json:"ref_block_prefix"`
				MaxNetUsageWords   int           `json:"max_net_usage_words"`
				MaxCPUUsageMs      int           `json:"max_cpu_usage_ms"`
				DelaySec           int           `json:"delay_sec"`
				ContextFreeActions []interface{} `json:"context_free_actions"`
				Actions            []struct {
					Account       string `json:"account"`
					Name          string `json:"name"`
					Authorization []struct {
						Actor      string `json:"actor"`
						Permission string `json:"permission"`
					} `json:"authorization"`
					Data string `json:"data"`
				} `json:"actions"`
				TransactionExtensions []interface{} `json:"transaction_extensions"`
				Signatures            []string      `json:"signatures"`
				ContextFreeData       []interface{} `json:"context_free_data"`
			}



			var sign_respon_0 sign_respon
			json.Unmarshal(body,&sign_respon_0)
			fmt.Println("获取签名Signatures：",sign_respon_0.Signatures)


			return sign_respon_0.Signatures,Actions_push,timestamp,RefBlockNum,RefBlockPrefix
			}else {
			fmt.Println("解锁失败")
		}


	}else if method=="delegatebw"{
		var wallet_name =configMap["wallet_name"]
		var wallet_password =configMap["wallet_password"]
		data_delegatebw:=delegatebwToBin.JsonToBin(account_name,buyCpuAmount,buyNetAmount)

		fmt.Println("获取编码后的data数据：",data_delegatebw)



		Actions_delegatebw:=Actions{
			Account:"eosio",
			Name:"delegatebw",
			Authorization:[]interface{}{Authorization_0},
			Data:data_delegatebw,
		}


		//Actions_1,_:=json.Marshal(Actions_0)

		sign_0:=sign{
			RefBlockNum:RefBlockNum,
			RefBlockPrefix:RefBlockPrefix,
			Expiration:timestamp,
			Actions:[]interface{}{Actions_delegatebw},
		}

		Actions_push:=[]interface{}{Actions_delegatebw}



		sign_transaction_1:=[]interface{}{sign_0,[]string{public_key},chain_id}

		sign_transaction_2,_:=json.Marshal(sign_transaction_1)

		sign_transaction :=string(sign_transaction_2)


		// fmt.Println(reflect.TypeOf(sign_0))


		isTrue:=lock.Unlock(wallet_name,wallet_password)
		if isTrue {
			body :=httpPost.HttpPost(sign_transaction,"wallet","sign_transaction")
			fmt.Println("sign_transaction签名返回数据：",string(body))


			type sign_respon struct {
				Expiration         string        `json:"expiration"`
				RefBlockNum        int           `json:"ref_block_num"`
				RefBlockPrefix     int           `json:"ref_block_prefix"`
				MaxNetUsageWords   int           `json:"max_net_usage_words"`
				MaxCPUUsageMs      int           `json:"max_cpu_usage_ms"`
				DelaySec           int           `json:"delay_sec"`
				ContextFreeActions []interface{} `json:"context_free_actions"`
				Actions            []struct {
					Account       string `json:"account"`
					Name          string `json:"name"`
					Authorization []struct {
						Actor      string `json:"actor"`
						Permission string `json:"permission"`
					} `json:"authorization"`
					Data string `json:"data"`
				} `json:"actions"`
				TransactionExtensions []interface{} `json:"transaction_extensions"`
				Signatures            []string      `json:"signatures"`
				ContextFreeData       []interface{} `json:"context_free_data"`
			}



			var sign_respon_0 sign_respon
			json.Unmarshal(body,&sign_respon_0)
			fmt.Println("获取签名Signatures：",sign_respon_0.Signatures)


			return sign_respon_0.Signatures,Actions_push,timestamp,RefBlockNum,RefBlockPrefix
		}else {
			fmt.Println("解锁失败")
		}


	}else {
		fmt.Println(method," 参数错误")
	}

   return  []string{"false"},[]interface{}{2222},"",0,0
	//return Actions_999
}