package main

import (
	"context"
	"crypto/ecdsa"
	"eth/foodchain"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

type Transaction struct {
	Id     *big.Int
	Amount *big.Int
}

func main() {

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("f62455b7c7e5b43b96ba88f5f2db0fe008802c7d15a247e40d3debe02a842a85")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance : ", balance)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress("0xA5958b4baED8379099fC8e987C92C0BcFB0d24C4")
	instance, err := foodchain.NewFoodchain(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	id := big.NewInt(1)
	amount := big.NewInt(100)

	tx, err := instance.AccountDeposit(auth, id, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s \n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	id, name, address, phoneNumber, balance, status, createdTime, err := instance.AccountInfo(&bind.CallOpts{}, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name, address, phoneNumber, balance, status, createdTime) // "bar"

	// read transaction
	// blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	myAbi, err := abi.JSON(strings.NewReader(foodchain.FoodchainMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println("Hash : ", tx.Hash().Hex())
		fmt.Println("Value : ", tx.Value().String())
		fmt.Println("Gas : ", tx.Gas())
		fmt.Println("Gas price : ", tx.GasPrice().Uint64())
		fmt.Println("Nonce : ", tx.Nonce())
		fmt.Println("Data :", tx.Data())
		fmt.Println("To : ", tx.To().Hex())
		var data Transaction

		err = myAbi.UnpackIntoInterface(data, "AccountDeposit", tx.Data())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Result : ", data)

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), big.NewInt(params.InitialBaseFee)); err == nil {
			fmt.Println("Message from:", msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
	}

}
