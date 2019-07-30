package main

import (
	"eosApi/server"
	"fmt"
)

func main()  {
	//参数栗子：输入钱包账户格式为："\"钱包名称"\"
	 wallet_name := "\"dacb_test11422\""

	 //return返回值  walletName和walletPassword 为string类型
     walletName,walletPassword:=server.CreateWallet(wallet_name)
     fmt.Println(walletName,walletPassword)
}




