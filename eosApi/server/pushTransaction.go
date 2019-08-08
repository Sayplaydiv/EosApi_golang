package server

import (
	"encoding/json"
	"eosApi/config"
	"eosApi/httpPost"
	"eosApi/server/lock"
	"fmt"
)

func PushTransaction(method string,account_name string,owner_public_key string,active_public_key1 string,active_public_key2 string,buyRamAmount string,buyCpuAmount string,buyNetAmount string)(string,int) {


	var configMap = config.InitConfig("src/eosApi/config/config.ini")
	var wallet_name =configMap["wallet_name"]


	type PushTransaction struct {
		Compression string `json:"compression"`
		Transaction interface{} `json:"transaction"`
		Signatures []string `json:"signatures"`
	}

	type Transaction struct {
		Expiration         string        `json:"expiration"`
		RefBlockNum        int           `json:"ref_block_num"`
		RefBlockPrefix     int         `json:"ref_block_prefix"`
		ContextFreeActions []interface{} `json:"context_free_actions"`
		Actions            []interface{} `json:"actions"`
		TransactionExtensions []interface{} `json:"transaction_extensions"`
	}

	Signatures,Actions_push,timestamp,RefBlockNum,RefBlockPrefix:=SignTransaction(method,account_name ,owner_public_key,active_public_key1,active_public_key2,buyRamAmount,buyCpuAmount,buyNetAmount)

	if Signatures[0]!="false"{

		//是否获取到信息
		/*

		fmt.Println("Signatures:",Signatures)
		fmt.Println("Actions_push:",Actions_push)
		fmt.Println("timestamp:",timestamp)
		fmt.Println("RefBlockNum:",RefBlockNum)
		fmt.Println("RefBlockPrefix:",RefBlockPrefix)

		*/

		Transaction_0:=Transaction{
			Expiration:timestamp,
			RefBlockNum:RefBlockNum,
			RefBlockPrefix:RefBlockPrefix,
			ContextFreeActions:[]interface{}{},
			Actions:Actions_push,
			TransactionExtensions:[]interface{}{},
		}

		PushTransaction_0 :=PushTransaction{
			Compression:"none",
			Transaction:Transaction_0,
			Signatures:Signatures,
		}

		PushTransaction_1,_:=json.Marshal(PushTransaction_0)
		PushTransaction_2:=string(PushTransaction_1)
		fmt.Println("PushTransaction请求数据：",PushTransaction_2)

		body:=httpPost.HttpPost(PushTransaction_2,"chain","push_transaction")

		if  method=="newaccount"{
			type Push_respon struct {
				TransactionID string `json:"transaction_id"`
				Processed     struct {
					ID              string      `json:"id"`
					BlockNum        int         `json:"block_num"`
					BlockTime       string      `json:"block_time"`
					ProducerBlockID interface{} `json:"producer_block_id"`
					Receipt         struct {
						Status        string `json:"status"`
						CPUUsageUs    int    `json:"cpu_usage_us"`
						NetUsageWords int    `json:"net_usage_words"`
					} `json:"receipt"`
					Elapsed      int  `json:"elapsed"`
					NetUsage     int  `json:"net_usage"`
					Scheduled    bool `json:"scheduled"`
					ActionTraces []struct {
						ActionOrdinal                          int `json:"action_ordinal"`
						CreatorActionOrdinal                   int `json:"creator_action_ordinal"`
						ClosestUnnotifiedAncestorActionOrdinal int `json:"closest_unnotified_ancestor_action_ordinal"`
						Receipt                                struct {
							Receiver       string          `json:"receiver"`
							ActDigest      string          `json:"act_digest"`
							GlobalSequence int             `json:"global_sequence"`
							RecvSequence   int             `json:"recv_sequence"`
							AuthSequence   [][]interface{} `json:"auth_sequence"`
							CodeSequence   int             `json:"code_sequence"`
							AbiSequence    int             `json:"abi_sequence"`
						} `json:"receipt"`
						Receiver string `json:"receiver"`
						Act      struct {
							Account       string `json:"account"`
							Name          string `json:"name"`
							Authorization []struct {
								Actor      string `json:"actor"`
								Permission string `json:"permission"`
							} `json:"authorization"`
							Data struct {
								Creator string `json:"creator"`
								Name    string `json:"name"`
								Owner   struct {
									Threshold int `json:"threshold"`
									Keys      []struct {
										Key    string `json:"key"`
										Weight int    `json:"weight"`
									} `json:"keys"`
									Accounts []interface{} `json:"accounts"`
									Waits    []interface{} `json:"waits"`
								} `json:"owner"`
								Active struct {
									Threshold int `json:"threshold"`
									Keys      []struct {
										Key    string `json:"key"`
										Weight int    `json:"weight"`
									} `json:"keys"`
									Accounts []interface{} `json:"accounts"`
									Waits    []interface{} `json:"waits"`
								} `json:"active"`
							} `json:"data"`
							HexData string `json:"hex_data"`
						} `json:"act"`
						ContextFree      bool        `json:"context_free"`
						Elapsed          int         `json:"elapsed"`
						Console          string      `json:"console"`
						TrxID            string      `json:"trx_id"`
						BlockNum         int         `json:"block_num"`
						BlockTime        string      `json:"block_time"`
						ProducerBlockID  interface{} `json:"producer_block_id"`
						AccountRAMDeltas []struct {
							Account string `json:"account"`
							Delta   int    `json:"delta"`
						} `json:"account_ram_deltas"`
						Except       interface{}   `json:"except"`
						ErrorCode    interface{}   `json:"error_code"`
						InlineTraces []interface{} `json:"inline_traces"`
					} `json:"action_traces"`
					AccountRAMDelta interface{} `json:"account_ram_delta"`
					Except          interface{} `json:"except"`
					ErrorCode       interface{} `json:"error_code"`
				} `json:"processed"`
			}

			fmt.Println("创建账户返回信息：",string(body))
			var  Push_respon_0  Push_respon
			json.Unmarshal(body,&Push_respon_0)
			transaction_id:=Push_respon_0.TransactionID
			block_num:=Push_respon_0.Processed.BlockNum
			account_name:=Push_respon_0.Processed.ActionTraces[0].Act.Data.Name
			owner_publickey:=Push_respon_0.Processed.ActionTraces[0].Act.Data.Owner.Keys[0].Key


			fmt.Println("交易hash为：",transaction_id)
			fmt.Println("交易所在区块为：",block_num)
			fmt.Println("创建的账户名称为：",account_name)
			fmt.Println("账户的owner权限公钥为：",owner_publickey)


			fmt.Println("提交交易返回信息为：",string(body))


			isTrue:=lock.Lock(wallet_name)
			if isTrue==false{
				fmt.Println("部署账户加锁失败")
			}

			return transaction_id,block_num

		}else if method=="buyram" {
			type Push_respon struct {
				TransactionID string `json:"transaction_id"`
				Processed     struct {
					ID              string      `json:"id"`
					BlockNum        int         `json:"block_num"`
					BlockTime       string      `json:"block_time"`
					ProducerBlockID interface{} `json:"producer_block_id"`
					Receipt         struct {
						Status        string `json:"status"`
						CPUUsageUs    int    `json:"cpu_usage_us"`
						NetUsageWords int    `json:"net_usage_words"`
					} `json:"receipt"`
					Elapsed      int  `json:"elapsed"`
					NetUsage     int  `json:"net_usage"`
					Scheduled    bool `json:"scheduled"`
					ActionTraces []struct {
						ActionOrdinal                          int `json:"action_ordinal"`
						CreatorActionOrdinal                   int `json:"creator_action_ordinal"`
						ClosestUnnotifiedAncestorActionOrdinal int `json:"closest_unnotified_ancestor_action_ordinal"`
						Receipt                                struct {
							Receiver       string          `json:"receiver"`
							ActDigest      string          `json:"act_digest"`
							GlobalSequence int             `json:"global_sequence"`
							RecvSequence   int             `json:"recv_sequence"`
							AuthSequence   [][]interface{} `json:"auth_sequence"`
							CodeSequence   int             `json:"code_sequence"`
							AbiSequence    int             `json:"abi_sequence"`
						} `json:"receipt"`
						Receiver string `json:"receiver"`
						Act      struct {
							Account       string `json:"account"`
							Name          string `json:"name"`
							Authorization []struct {
								Actor      string `json:"actor"`
								Permission string `json:"permission"`
							} `json:"authorization"`
							Data struct {
								Payer    string `json:"payer"`
								Receiver string `json:"receiver"`
								Quant    string `json:"quant"`
							} `json:"data"`
							HexData string `json:"hex_data"`
						} `json:"act"`
						ContextFree      bool          `json:"context_free"`
						Elapsed          int           `json:"elapsed"`
						Console          string        `json:"console"`
						TrxID            string        `json:"trx_id"`
						BlockNum         int           `json:"block_num"`
						BlockTime        string        `json:"block_time"`
						ProducerBlockID  interface{}   `json:"producer_block_id"`
						AccountRAMDeltas []interface{} `json:"account_ram_deltas"`
						Except           interface{}   `json:"except"`
						ErrorCode        interface{}   `json:"error_code"`
						InlineTraces     []struct {
							ActionOrdinal                          int `json:"action_ordinal"`
							CreatorActionOrdinal                   int `json:"creator_action_ordinal"`
							ClosestUnnotifiedAncestorActionOrdinal int `json:"closest_unnotified_ancestor_action_ordinal"`
							Receipt                                struct {
								Receiver       string          `json:"receiver"`
								ActDigest      string          `json:"act_digest"`
								GlobalSequence int             `json:"global_sequence"`
								RecvSequence   int             `json:"recv_sequence"`
								AuthSequence   [][]interface{} `json:"auth_sequence"`
								CodeSequence   int             `json:"code_sequence"`
								AbiSequence    int             `json:"abi_sequence"`
							} `json:"receipt"`
							Receiver string `json:"receiver"`
							Act      struct {
								Account       string `json:"account"`
								Name          string `json:"name"`
								Authorization []struct {
									Actor      string `json:"actor"`
									Permission string `json:"permission"`
								} `json:"authorization"`
								Data struct {
									From     string `json:"from"`
									To       string `json:"to"`
									Quantity string `json:"quantity"`
									Memo     string `json:"memo"`
								} `json:"data"`
								HexData string `json:"hex_data"`
							} `json:"act"`
							ContextFree      bool          `json:"context_free"`
							Elapsed          int           `json:"elapsed"`
							Console          string        `json:"console"`
							TrxID            string        `json:"trx_id"`
							BlockNum         int           `json:"block_num"`
							BlockTime        string        `json:"block_time"`
							ProducerBlockID  interface{}   `json:"producer_block_id"`
							AccountRAMDeltas []interface{} `json:"account_ram_deltas"`
							Except           interface{}   `json:"except"`
							ErrorCode        interface{}   `json:"error_code"`
							InlineTraces     []struct {
								ActionOrdinal                          int `json:"action_ordinal"`
								CreatorActionOrdinal                   int `json:"creator_action_ordinal"`
								ClosestUnnotifiedAncestorActionOrdinal int `json:"closest_unnotified_ancestor_action_ordinal"`
								Receipt                                struct {
									Receiver       string          `json:"receiver"`
									ActDigest      string          `json:"act_digest"`
									GlobalSequence int             `json:"global_sequence"`
									RecvSequence   int             `json:"recv_sequence"`
									AuthSequence   [][]interface{} `json:"auth_sequence"`
									CodeSequence   int             `json:"code_sequence"`
									AbiSequence    int             `json:"abi_sequence"`
								} `json:"receipt"`
								Receiver string `json:"receiver"`
								Act      struct {
									Account       string `json:"account"`
									Name          string `json:"name"`
									Authorization []struct {
										Actor      string `json:"actor"`
										Permission string `json:"permission"`
									} `json:"authorization"`
									Data struct {
										From     string `json:"from"`
										To       string `json:"to"`
										Quantity string `json:"quantity"`
										Memo     string `json:"memo"`
									} `json:"data"`
									HexData string `json:"hex_data"`
								} `json:"act"`
								ContextFree      bool          `json:"context_free"`
								Elapsed          int           `json:"elapsed"`
								Console          string        `json:"console"`
								TrxID            string        `json:"trx_id"`
								BlockNum         int           `json:"block_num"`
								BlockTime        string        `json:"block_time"`
								ProducerBlockID  interface{}   `json:"producer_block_id"`
								AccountRAMDeltas []interface{} `json:"account_ram_deltas"`
								Except           interface{}   `json:"except"`
								ErrorCode        interface{}   `json:"error_code"`
								InlineTraces     []interface{} `json:"inline_traces"`
							} `json:"inline_traces"`
						} `json:"inline_traces"`
					} `json:"action_traces"`
					AccountRAMDelta interface{} `json:"account_ram_delta"`
					Except          interface{} `json:"except"`
					ErrorCode       interface{} `json:"error_code"`
				} `json:"processed"`
			}


			fmt.Println("购买ram返回数据：",string(body))
			var  Push_respon_0  Push_respon
			json.Unmarshal(body,&Push_respon_0)
			transaction_id:=Push_respon_0.TransactionID
			block_num:=Push_respon_0.Processed.BlockNum

			fmt.Println("交易hash为：",transaction_id)
			fmt.Println("交易所在区块为：",block_num)


			isTrue:=lock.Lock(wallet_name)
			if isTrue==false{
				fmt.Println("部署账户加锁失败")
			}

			return  transaction_id,block_num
		}else if method=="delegatebw"{
			type Push_respon struct {
				TransactionID string `json:"transaction_id"`
				Processed     struct {
					ID              string      `json:"id"`
					BlockNum        int         `json:"block_num"`
					BlockTime       string      `json:"block_time"`
					ProducerBlockID interface{} `json:"producer_block_id"`
					Receipt         struct {
						Status        string `json:"status"`
						CPUUsageUs    int    `json:"cpu_usage_us"`
						NetUsageWords int    `json:"net_usage_words"`
					} `json:"receipt"`
					Elapsed      int  `json:"elapsed"`
					NetUsage     int  `json:"net_usage"`
					Scheduled    bool `json:"scheduled"`
					ActionTraces []struct {
						ActionOrdinal                          int `json:"action_ordinal"`
						CreatorActionOrdinal                   int `json:"creator_action_ordinal"`
						ClosestUnnotifiedAncestorActionOrdinal int `json:"closest_unnotified_ancestor_action_ordinal"`
						Receipt                                struct {
							Receiver       string          `json:"receiver"`
							ActDigest      string          `json:"act_digest"`
							GlobalSequence int             `json:"global_sequence"`
							RecvSequence   int             `json:"recv_sequence"`
							AuthSequence   [][]interface{} `json:"auth_sequence"`
							CodeSequence   int             `json:"code_sequence"`
							AbiSequence    int             `json:"abi_sequence"`
						} `json:"receipt"`
						Receiver string `json:"receiver"`
						Act      struct {
							Account       string `json:"account"`
							Name          string `json:"name"`
							Authorization []struct {
								Actor      string `json:"actor"`
								Permission string `json:"permission"`
							} `json:"authorization"`
							Data struct {
								Payer    string `json:"payer"`
								Receiver string `json:"receiver"`
								Quant    string `json:"quant"`
							} `json:"data"`
							HexData string `json:"hex_data"`
						} `json:"act"`
						ContextFree      bool          `json:"context_free"`
						Elapsed          int           `json:"elapsed"`
						Console          string        `json:"console"`
						TrxID            string        `json:"trx_id"`
						BlockNum         int           `json:"block_num"`
						BlockTime        string        `json:"block_time"`
						ProducerBlockID  interface{}   `json:"producer_block_id"`
						AccountRAMDeltas []interface{} `json:"account_ram_deltas"`
						Except           interface{}   `json:"except"`
						ErrorCode        interface{}   `json:"error_code"`
						InlineTraces     []struct {
							ActionOrdinal                          int `json:"action_ordinal"`
							CreatorActionOrdinal                   int `json:"creator_action_ordinal"`
							ClosestUnnotifiedAncestorActionOrdinal int `json:"closest_unnotified_ancestor_action_ordinal"`
							Receipt                                struct {
								Receiver       string          `json:"receiver"`
								ActDigest      string          `json:"act_digest"`
								GlobalSequence int             `json:"global_sequence"`
								RecvSequence   int             `json:"recv_sequence"`
								AuthSequence   [][]interface{} `json:"auth_sequence"`
								CodeSequence   int             `json:"code_sequence"`
								AbiSequence    int             `json:"abi_sequence"`
							} `json:"receipt"`
							Receiver string `json:"receiver"`
							Act      struct {
								Account       string `json:"account"`
								Name          string `json:"name"`
								Authorization []struct {
									Actor      string `json:"actor"`
									Permission string `json:"permission"`
								} `json:"authorization"`
								Data struct {
									From     string `json:"from"`
									To       string `json:"to"`
									Quantity string `json:"quantity"`
									Memo     string `json:"memo"`
								} `json:"data"`
								HexData string `json:"hex_data"`
							} `json:"act"`
							ContextFree      bool          `json:"context_free"`
							Elapsed          int           `json:"elapsed"`
							Console          string        `json:"console"`
							TrxID            string        `json:"trx_id"`
							BlockNum         int           `json:"block_num"`
							BlockTime        string        `json:"block_time"`
							ProducerBlockID  interface{}   `json:"producer_block_id"`
							AccountRAMDeltas []interface{} `json:"account_ram_deltas"`
							Except           interface{}   `json:"except"`
							ErrorCode        interface{}   `json:"error_code"`
							InlineTraces     []struct {
								ActionOrdinal                          int `json:"action_ordinal"`
								CreatorActionOrdinal                   int `json:"creator_action_ordinal"`
								ClosestUnnotifiedAncestorActionOrdinal int `json:"closest_unnotified_ancestor_action_ordinal"`
								Receipt                                struct {
									Receiver       string          `json:"receiver"`
									ActDigest      string          `json:"act_digest"`
									GlobalSequence int             `json:"global_sequence"`
									RecvSequence   int             `json:"recv_sequence"`
									AuthSequence   [][]interface{} `json:"auth_sequence"`
									CodeSequence   int             `json:"code_sequence"`
									AbiSequence    int             `json:"abi_sequence"`
								} `json:"receipt"`
								Receiver string `json:"receiver"`
								Act      struct {
									Account       string `json:"account"`
									Name          string `json:"name"`
									Authorization []struct {
										Actor      string `json:"actor"`
										Permission string `json:"permission"`
									} `json:"authorization"`
									Data struct {
										From     string `json:"from"`
										To       string `json:"to"`
										Quantity string `json:"quantity"`
										Memo     string `json:"memo"`
									} `json:"data"`
									HexData string `json:"hex_data"`
								} `json:"act"`
								ContextFree      bool          `json:"context_free"`
								Elapsed          int           `json:"elapsed"`
								Console          string        `json:"console"`
								TrxID            string        `json:"trx_id"`
								BlockNum         int           `json:"block_num"`
								BlockTime        string        `json:"block_time"`
								ProducerBlockID  interface{}   `json:"producer_block_id"`
								AccountRAMDeltas []interface{} `json:"account_ram_deltas"`
								Except           interface{}   `json:"except"`
								ErrorCode        interface{}   `json:"error_code"`
								InlineTraces     []interface{} `json:"inline_traces"`
							} `json:"inline_traces"`
						} `json:"inline_traces"`
					} `json:"action_traces"`
					AccountRAMDelta interface{} `json:"account_ram_delta"`
					Except          interface{} `json:"except"`
					ErrorCode       interface{} `json:"error_code"`
				} `json:"processed"`
			}


			fmt.Println("抵押cpu和net返回数据：",string(body))
			var  Push_respon_0  Push_respon
			json.Unmarshal(body,&Push_respon_0)
			transaction_id:=Push_respon_0.TransactionID
			block_num:=Push_respon_0.Processed.BlockNum

			fmt.Println("交易hash为：",transaction_id)
			fmt.Println("交易所在区块为：",block_num)


			isTrue:=lock.Lock(wallet_name)
			if isTrue==false{
				fmt.Println("部署账户加锁失败")
			}
			return  transaction_id,block_num
		}
	}else {
		fmt.Println("签名失败,请检查网络及密码是否正确")
	}
        return "",0
}
