package main

import (
	"eosApi/server"
	"fmt"
)

func main()  {


	//使用的方法
	var method="buyram"

	//接收ram的账户
	var account_name="eostest22322"

	//要购买ram的数量
	var buyRamAmount="1.0000 EOS"

	transaction_id,block_num:=server.PushTransaction(method,account_name,"","","",buyRamAmount,"","")
	fmt.Println(transaction_id,block_num)
}