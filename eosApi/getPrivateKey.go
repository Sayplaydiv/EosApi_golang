package main

import (
	"eosApi/server"
	"fmt"
)

func main()  {

	//参数栗子
	walletName:="dacbtest111"
	walletPassword:="PW5J4RParpAxZy7N5y1N4VuCsktL9Y7iWumZgfC5oYVhjAnMoJzXN"

	//return返回数据list_keys格式为二维数组，[][]string
	list_keys:=server.GetPrivateKeys(walletName,walletPassword)
	fmt.Println("公私钥对列表：",list_keys)
}
