package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	
	//algo package
	"github.com/algorand/go-algorand-sdk/future"
    "github.com/algorand/go-algorand-sdk/client/v2/algod"
    "github.com/algorand/go-algorand-sdk/client/v2/common"
    "context"
    json "encoding/json"
)


type SignedTransaction struct {
	Txn []uint8 `json:"signedTXN"`
	TxnID string `json:"txID"`
}

func broadcast (c *gin.Context) {

	const algodAddress = "https://testnet-algorand.api.purestake.io/ps2"
  	const psTokenKey = "X-API-Key"
  	const psToken = ""

  	commonClient, err := common.MakeClient(algodAddress, psTokenKey, psToken)
	if err != nil {
		fmt.Printf("failed to make common client: %s\n", err)
		return
	}
	algodClient := (*algod.Client)(commonClient)


	fmt.Println("Status")
  	nodeStatus, err := algodClient.Status().Do(context.Background())
  	if err != nil {
    	fmt.Printf("error getting algod status: %s\n", err)
    	return
  	}

  	fmt.Printf("algod last round: %d\n", nodeStatus.LastRound)

	var stxn SignedTransaction
	if err := c.BindJSON(&stxn); err != nil {
		fmt.Println(stxn.Txn)
		c.JSON(400, gin.H{"error": "Transaction error"})
		return
	}

	c.JSON(200, gin.H{"transaction": stxn.Txn})
	fmt.Println(stxn.Txn)


	sendResponse, err := algodClient.SendRawTransaction(stxn.Txn).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)

	confirmedTxn, err := future.WaitForConfirmation(algodClient, stxn.TxnID, 4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", stxn.TxnID)
		return
	}

	txnJSON, err := json.MarshalIndent(confirmedTxn.Transaction.Txn, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshal txn data: %s\n", err)
	}

	fmt.Printf("Transaction information: %s\n", txnJSON)
	fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))
	fmt.Printf("Amount sent: %d microAlgos\n", confirmedTxn.Transaction.Txn.Amount)
	fmt.Printf("Fee: %d microAlgos\n", confirmedTxn.Transaction.Txn.Fee)    
	fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))
	

}


func main () {
	router := gin.Default()
	router.POST("/broadcast", broadcast)
	router.Run(":500")
}