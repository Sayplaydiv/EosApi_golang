package main

import (
	"eosApi/server"
	"fmt"
)

func main()  {


	//接收cpu和net的账户
	var account_name="eostest22322"

	//给新账户抵押cpu的数量
	var buyCpuAmount="1.0000 EOS"

	//给新账户抵押net的数量
	var buyNetAmount="1.0000 EOS"

	//使用的方法
	var method="delegatebw"

	//交易transaction_id为string,交易所在区块block_num为
	transaction_id,block_num:=server.PushTransaction(method,account_name,"","","","",buyCpuAmount,buyNetAmount)
	fmt.Println(transaction_id,block_num)

}