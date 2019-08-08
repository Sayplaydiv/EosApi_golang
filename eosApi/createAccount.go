package main

import (
	"eosApi/server"
	"fmt"
)

func main()  {

	//账户名称必须为12位1-5，a-z

	//例如：
	//创建的账户名称
	var create_account_name="eostest33222"

	//owner公钥
	var owner_publickey="EOS7JwTdbdFPF8MQ5XGityebKd6V19RrbTF9JSs2QYdkKv9FXDmJ6"

	//active公钥1
	var active_publickey1="EOS4y2t6dScxQQgD9xZtdB25ZAs3DdU2f84L9txusXH9L3SHFtbna"

	//active公钥2
	var active_publickey2="EOS7krT1L3btH74cErR2GeRxU9SRUdMwoPUSNfsPbgL3VVemNXL1F"

	//给新账户购买ram的数量
	var buyRamAmount="1.0000 EOS"

	//给新账户抵押cpu的数量
	var buyCpuAmount="1.0000 EOS"

	//给新账户抵押net的数量
	var buyNetAmount="1.0000 EOS"

	//使用的方法
	var method="newaccount"


	//创建成功，return返回值：block_num交易所在块为int类型，transaction_id为string类型
	transaction_id,block_num:=server.PushTransaction(method,create_account_name,owner_publickey,active_publickey1,active_publickey2,buyRamAmount,buyCpuAmount,buyNetAmount)
	fmt.Println(transaction_id,block_num)
}