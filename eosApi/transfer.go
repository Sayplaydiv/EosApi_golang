package main

import (
	"eosApi/server"
	"fmt"
)

func main() {
	//参数栗子

	//出账用户名称
	from_address:="dacbtest1111"

	//收账用户名称
	to_address:="dacbeostest2"

	//转帐金额
	quantity:="20.0000 EOS"

	//memo备注信息
	memo:="这里是备注"

	//出账账户所在的钱包名称（用于签名）：
	wallet_name:="dacbtest111"

	//出账账户所在的密码名称（用于签名）：
	wallet_password:="PW5J4RParpAxZy7N5y1N4VuCsktL9Y7iWumZgfC5oYVhjAnMoJzXN"

	//转帐的active权限（我一般用多重签名，两把密钥对来控制一个权限，如果active只有一个active_publicKey2置为空即可如： ""）
	active_publicKey1:="EOS72QpGhNGtVi6QALhiQQhBZ2VvvjWAvcQ2jDcGLQ7tLPidD7KPJ"
	active_publicKey2:="EOS7xs9YJhZuPaYU2NuYZEbov1xorEEJvomKLyRwiVDCAmqpuALm3"



	//return返回值交易hash的transaction_id为string类型，block_num为int类型
	transaction_id,block_num:=server.PushTransfer(from_address,to_address,quantity,memo,wallet_name,wallet_password,active_publicKey1,active_publicKey2)
	fmt.Println(transaction_id,block_num)
	
}
