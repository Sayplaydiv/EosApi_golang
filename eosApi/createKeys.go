package main

import (
	"eosApi/server"
	"fmt"
)

func main()  {

	//参数栗子
	walletName:="dacbtest111"
	walletPassword:="PW5J4RParpAxZy7N5y1N4VuCsktL9Y7iWumZgfC5oYVhjAnMoJzXN"


	//return返回值Public_key为string类型
	Public_key:=server.CreateKeys(walletName,walletPassword)
	fmt.Println("创建的密钥对公钥为：",Public_key)
}
