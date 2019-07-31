package main

import (
	"eosApi/server"
	"fmt"
)

func main() {
	//参数栗子：


	//需要查询的块高度
	blockNum:=41897731


	//return返回为json字符串，可根据需求解析相应的结构体
	body_json:=server.GetBlockInfo(blockNum)
	fmt.Println(body_json)
}
