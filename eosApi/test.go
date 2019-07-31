package main

func main(){

	//购买ram编码
	//data_buyRam:=buyRamToBin.JsonToBin("dacbtest2222","1.0000 EOS")


	//创建账户编码
	//data_newAccount:=newAccountToBin.JsonToBin("dacbtest3333","EOS72QpGhNGtVi6QALhiQQhBZ2VvvjWAvcQ2jDcGLQ7tLPidD7KPJ","EOS7xs9YJhZuPaYU2NuYZEbov1xorEEJvomKLyRwiVDCAmqpuALm3","EOS7gb3wkq8RgV19zjNBuuVA2PNd98qPPFZ4Eox67qt7KZmAiFdq8")


    //抵押cpu和net
	//data_delegatebw:=delegatebwToBin.JsonToBin("dacbtest3333","1.0000 EOS","1.0000 EOS")



	//获取区块高度和id
	//server.GetInfo()


	//获取指定高度区块信息
	//server.GetBlock()


	//签名
	/*

	Signatures,Actions_push,timestamp,RefBlockNum,RefBlockPrefix:=server.SignTransaction()
	fmt.Println("Signatures:",Signatures)
	fmt.Println("Actions_push:",Actions_push)
	fmt.Println("timestamp:",timestamp)
	fmt.Println("RefBlockNum:",RefBlockNum)
	fmt.Println("RefBlockPrefix:",RefBlockPrefix)

	*/


	//账户解锁
	/*

	isTrue:=lock.Unlock()
	if  isTrue{
		lock.Lock()
	}

	*/

	//提交交易
	//server.PushTransaction()

	//buyram
	//server.SignTransaction("buyram","eostest22322","","","","1.0000 EOS","","")

	//抵押cpu和net
	//server.SignTransaction("delegatebw","eostest22322","","","","","1.0000 EOS","1.0000 EOS")

	//transfer编码
	//transfer_data:=transfer.JsonToBin("dacbtest1111","eostest22322","1.0000 EOS","qq771254336")
	//fmt.Println(transfer_data)


	//5JJj9ijPhnauwAaY2vnVFooc5VVLetXct27vtaEaHt1GqE16ojr
	//5JGnraLC2XUy3dyt15EBjDQdy63gdt4n7BTPpcV68rbYewhvgza

	//转帐交易签名：
	/*
	Signatures,Actions_push,timestamp,RefBlockNum,RefBlockPrefix:=server.SignTransfer("transfer","dacbtest1111","eostest22322","10.0000 EOS","qq771254336","dacbtest111","PW5J4RParpAxZy7N5y1N4VuCsktL9Y7iWumZgfC5oYVhjAnMoJzXN","EOS72QpGhNGtVi6QALhiQQhBZ2VvvjWAvcQ2jDcGLQ7tLPidD7KPJ","EOS7xs9YJhZuPaYU2NuYZEbov1xorEEJvomKLyRwiVDCAmqpuALm3")
	fmt.Println("Signatures:",Signatures)
	fmt.Println("Actions_push:",Actions_push)
	fmt.Println("timestamp:",timestamp)
	fmt.Println("RefBlockNum:",RefBlockNum)
	fmt.Println("RefBlockPrefix:",RefBlockPrefix)
	*/

	//transaction_id,block_num:=server.PushTransfer("dacbtest1111","eostest22322","1.0000 EOS","qq771254336","dacbtest111","PW5J4RParpAxZy7N5y1N4VuCsktL9Y7iWumZgfC5oYVhjAnMoJzXN","EOS72QpGhNGtVi6QALhiQQhBZ2VvvjWAvcQ2jDcGLQ7tLPidD7KPJ","EOS7xs9YJhZuPaYU2NuYZEbov1xorEEJvomKLyRwiVDCAmqpuALm3")

	//fmt.Println(transaction_id,block_num)

}
