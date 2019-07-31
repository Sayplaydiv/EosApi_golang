package main

import (
	"eosApi/server"
	"fmt"
)

func main() {
	last_block:=server.GetInfo()
	fmt.Println("最新块高为：",last_block)
}
