package main

import (
	"eosApi/server"
	"fmt"
)

func main()  {

	//账户名称必须为12位1-5，a-z

	//例如：
	//创建的账户名称
	var create_account_name="dacbeostest2"

	//owner公钥
	//私钥：5JcfMKkwG52PXid7KgjqsUXdMVAzVc4SkhKa1YRLMTnHkL8u51s
	var owner_publickey="EOS5fvqGKW9jDRvRnboCuG2ErjHjKcRbqEixXhsEcSgTB88F6VEjm"

	//active公钥1
	//私钥：5K3SGFG97ERDnaeJfNpFXoFmQuKyjWC8LpQxjfJ9b37tc8jDV7A
	var active_publickey1="EOS5utxT48sHmbKeTSQFNLVeS5VvHWZ9nDpvtjtkkAVg6EGp41NR5"

	//active公钥2
	//var active_publickey2="EOS7krT1L3btH74cErR2GeRxU9SRUdMwoPUSNfsPbgL3VVemNXL1F"

	//给新账户购买ram的数量
	var buyRamAmount="1.0000 EOS"

	//给新账户抵押cpu的数量
	var buyCpuAmount="1.0000 EOS"

	//给新账户抵押net的数量
	var buyNetAmount="1.0000 EOS"

	//使用的方法
	var method="newaccount"


	//创建成功，return返回值：block_num交易所在块为int类型，transaction_id为string类型
	transaction_id,block_num:=server.PushTransaction(method,create_account_name,owner_publickey,active_publickey1,"",buyRamAmount,buyCpuAmount,buyNetAmount)
	fmt.Println(transaction_id,block_num)
}