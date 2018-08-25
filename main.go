package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(6213364)

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
		fmt.Printf("TX Value: %s\n", tx.Value().String())
		fmt.Printf("TX Gas: %d\n", tx.Gas())
		fmt.Printf("TX Gas Price: %d\n", tx.GasPrice().Uint64())
		fmt.Printf("TX Nonce: %d\n", tx.Nonce())
		fmt.Printf("TX Data: %v\n", tx.Data())
		fmt.Printf("TX To: %s\n", tx.To().Hex())

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Receipt Status: %d\n", receipt.Status)
		fmt.Println("---")
	}

	// Grab block by hash then iterate over transactions by index
	blockHash := common.HexToHash("0x2a875a424a5236d5ae8f3524c86f158abe63499ee6089ca07abd01e5b1257cb1")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
	}

	// Grab a transaction by it's individual hash
	txHash := common.HexToHash("0xae9c3776de9ed6bf0e025704bbeced567b428c78e00330b59c25fe45e7ef87a9")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
	fmt.Printf("Pending?: %v\n", isPending)
}
