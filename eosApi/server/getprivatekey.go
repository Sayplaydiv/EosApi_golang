package server

import (
	"encoding/json"
	"eosApi/httpPost"
	"fmt"
)

func GetPrivateKeys(walletName string,walletPassword string) ([][]string) {

	listkeys_0:=[]string{walletName,walletPassword}
	listkeys_1,_:=json.Marshal(listkeys_0)
	listkeys:=string(listkeys_1)
	body:=httpPost.HttpPost(listkeys,"wallet","list_keys")

	//fmt.Println(string(body))
	var list_keys [][]string
	json.Unmarshal(body,&list_keys)

	fmt.Println("钱包内生产公私钥对的数量为：",len(list_keys))
	for i:=0;i<len(list_keys) ;i++  {
		Public_key:=list_keys[i][0]
		Private_key:=list_keys[i][1]

		fmt.Println("第",i,"组公钥为：",Public_key)
		fmt.Println("第",i,"组私钥为：",Private_key)
	}
	return list_keys

}