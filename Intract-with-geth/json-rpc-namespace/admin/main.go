package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.DialContext(context.Background(), "http://127.0.0.1:8551")
	if err != nil {
		fmt.Println(err.Error() + ": Gak bisa connect")
	}
	defer client.Close()

	balance, _ := client.BalanceAt(context.Background(), common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"), big.NewInt(19725990))

	fmt.Println(balance.String())
}
