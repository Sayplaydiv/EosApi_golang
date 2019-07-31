package main

import (
	"eosApi/server"
	"fmt"
)

func main() {
	//参数栗子：

	//发行账户：EOS为eosio.token，其他代币根据添其发行账户的名称，例如BMC代币的发行账户为blockiotoken
	code:="eosio.token"
	code_0:="blockiotoken"

	//需要查询的账户名称
	account:="eostest22322"
	account_0:="junglefaucet"

	//需要查询的币种名称
    symbol:="EOS"
	symbol_0:="BMC"



	//return返回值为[]string数组
	balance:=server.GetBalance(code,account,symbol)
	balance_0:=server.GetBalance(code_0,account_0,symbol_0)

	fmt.Println("EOS余额为：",balance)
	fmt.Println("BMC余额为：",balance_0)

}
