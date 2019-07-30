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

    	
    ##节点接口：
    chain_url=http://13.52.165.98:8888
    
    ##钱包接口：
    wallet_url=http://13.52.165.98:8900
    
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