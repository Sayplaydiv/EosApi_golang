package server

import (
	"encoding/json"
	"eosApi/config"
	"eosApi/httpPost"
	"eosApi/server/lock"
	"eosApi/server/transfer"
	"fmt"
	"strings"
	"time"
)

func SignTransfer(method string,from_address string,to_address string,quantity string,memo string,wallet_name string,wallet_password string,active_publicKye1 string,active_publicKye2 string)([]string,[]interface{},string,int,int) {
	var configMap = config.InitConfig("src/eosApi/config/config.ini")
	var chain_id =configMap["chain_id"]


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
		Actor:from_address,
		Permission:"active",
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


	 if method=="transfer"{

		data_transfer:=transfer.JsonToBin(from_address,to_address,quantity,memo)

		fmt.Println("获取编码后的data数据:",data_transfer)
		//Authorization_1,_:=json.Marshal(Authorization_0)


		Actions_transfer:=Actions{
			Account:"eosio.token",
			Name:"transfer",
			Authorization:[]interface{}{Authorization_0},
			Data:data_transfer,
		}


		//Actions_1,_:=json.Marshal(Actions_0)

		sign_0:=sign{
			RefBlockNum:RefBlockNum,
			RefBlockPrefix:RefBlockPrefix,
			Expiration:timestamp,
			Actions:[]interface{}{Actions_transfer},
		}

		Actions_push:=[]interface{}{Actions_transfer}



		sign_transfer_1:=[]interface{}{sign_0,[]string{active_publicKye1,active_publicKye2},chain_id}

		sign_transfer_2,_:=json.Marshal(sign_transfer_1)

		sign_transfer :=string(sign_transfer_2)


		// fmt.Println(reflect.TypeOf(sign_0))


		isTrue:=lock.Unlock(wallet_name,wallet_password)
		if isTrue {
			body :=httpPost.HttpPost(sign_transfer,"wallet","sign_transaction")
			fmt.Println("sign_transfer签名返回数据：",string(body))


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
	
}
