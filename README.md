4
##EOSapi文档 golang

####1.创建钱包：执行入口createWallet.go


//参数栗子：输入钱包账户格式为："\"钱包名称"\"
```	
	 wallet_name := "\"eostest111\""

```
 //return返回值  walletName和walletPassword 为string类型
``` 
    walletName,walletPassword:=server.CreateWallet(wallet_name)
    fmt.Println(walletName,walletPassword)
```



####2.创建公私钥对：执行入口createKeys.go

//参数栗子
```	
	walletName:="eostest111"
	walletPassword:="PW5J4RParpAxZy7N5y1N4VuCsktL9Y7iWumZgfC5oYVhjAnMoJzXN"
```

//return返回值Public_key为string类型
```
	Public_key:=server.CreateKeys(walletName,walletPassword)
	fmt.Println("创建的密钥对公钥为：",Public_key)
```

####3.获取公私钥：执行入口getPrivateKey.go
//参数栗子
```
	walletName:="eostest111"
	walletPassword:="PW5J4RParpAxZy7N5y1N4VuCsktL9Y7iWumZgfC5oYVhjAnMoJzXN"
```
//return返回数据list_keys格式为二维数组，[][]string
```	
	list_keys:=server.GetPrivateKeys(walletName,walletPassword)
	fmt.Println("公私钥对列表：",list_keys)
```

####4.创建账户：执行入口createAccount.go
######账户名称必须为12位1-5，a-z
    
######栗子：
   //创建的账户名称
    	
   
   var create_account_name="eostest22322"
  
    
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
    	
####5.配置文件：

#####位置：~/config/config.ini
#####准备：一个用于创建，购买ram抵押等的其他账户的部署账户
    	
    ##节点接口：
    chain_url=http://13.xx.xx.xx:8888
    
    ##钱包接口：
    wallet_url=http://13.xx.xx.xx:8900
    
    ##测试网络id：e70aaab8997e1dfce58fbfac80cbbb8fecec7b99cf982a9444273cbc64c41473
    ##主网id： aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906
    chain_id=e70aaab8997e1dfce58fbfac80cbbb8fecec7b99cf982a9444273cbc64c41473
    
    ##钱包名称：
    wallet_name=eostest111
    
    ##钱包密码：
    wallet_password=PW5J4RParpAxZy7N5y1N4VuCsktL9Y7iWumZgfC5oYVhjAnMoJzXN
    
    ##部署账户的owner权限公钥（为了方便，实际开发可根据需要的权限用active或owner）：
    account_public_key=EOS71PMU3TUvx92YgUgYJNbVM66jFU3P9fSyU7ZaGuJhiKTAeJMHp
    
    ##部署账户名称：
    super_account=dacbtest1111	 	
    
    
    
####6.购买ram内存：执行入口buyRam.go
//参数栗子：

//使用的方法
```
	var method="buyram"
```
//接收ram的账户

```
	var account_name="eostest22322"
```

//要购买ram的数量

```
	var buyRamAmount="1.0000 EOS"
```

//购买成功，return返回值：block_num交易所在块为int类型，transaction_id为string类型

```
	transaction_id,block_num:=server.PushTransaction(method,account_name,"","","",buyRamAmount,"","")
	fmt.Println(transaction_id,block_num)    
```    

####7.抵押cpu和net：执行入口delegatebw.go
//参数栗子：

//接收cpu和net的账户
```
	var account_name="eostest22322"
```
//给新账户抵押cpu的数量
```
	var buyCpuAmount="1.0000 EOS"
```
//给新账户抵押net的数量
```	
	var buyNetAmount="1.0000 EOS"
```
//使用的方法
```	
	var method="delegatebw"
```
//抵押成功，交易transaction_id为string,交易所在区块block_num为
```	
	transaction_id,block_num:=server.PushTransaction(method,account_name,"","","","",buyCpuAmount,buyNetAmount)
	fmt.Println(transaction_id,block_num)
```

####8.转帐操作：执行入口transfer.go
//参数栗子

//出账用户名称

	from_address:="dacbtest1111"

//收账用户名称

	to_address:="eostest22322"

//转帐金额

	quantity:="1.0000 EOS"

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

####9.查询余额：执行入口getBlance.go
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
	
####10,查询最新块高：执行入口getInfo.go

	last_block:=server.GetInfo()
	fmt.Println("最新块高为：",last_block)	
	
####11.查询指定块高的区块信息：执行入口getBlcokInfo.go
//参数栗子：

需要查询的块高度

	blockNum:=41897731

//return返回为json字符串，可根据需求解析相应的结构体

	body_json:=server.GetBlockInfo(blockNum)
	fmt.Println(body_json)

}	
	
	
