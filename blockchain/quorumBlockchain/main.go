package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-simple-storage-fcc/api"
	"math/big"
)

var gateway = "http://127.0.0.1:22000"
var accountPrivateKey = "9fa2a112909c080d807dbb6efe4ef21c257b8880e8ff7167161fa919219f7c43"

func main() {
	_, client, fromAddress, privateKey, chainID, simpleStorageApi := contractDeploy()
	createWallet(client, fromAddress, privateKey, chainID, simpleStorageApi, "First", 30000)
	infoWallet(simpleStorageApi, "First")
	createWallet(client, fromAddress, privateKey, chainID, simpleStorageApi, "Two", 10000)
	infoWallet(simpleStorageApi, "Two")
	fmt.Println("transfer money:")
	transferMoney(client, fromAddress, privateKey, chainID, simpleStorageApi, "First", "Two", 3000)
	infoWallet(simpleStorageApi, "First")
	infoWallet(simpleStorageApi, "Two")
}

// GetNextTransaction returns the next transaction in the pending transaction queue
func getNextTransaction(client *ethclient.Client, fromAddress common.Address, privateKey *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	// nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	// sign the transaction
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)             // in wei
	auth.GasLimit = uint64(574041)         // in units
	auth.GasPrice = big.NewInt(1000000000) // in wei

	return auth, nil
}

func contractDeploy() (*bind.TransactOpts, *ethclient.Client, common.Address, *ecdsa.PrivateKey, *big.Int, *api.Api) {
	// connect to blockchain network
	client, err := ethclient.Dial(gateway)
	if err != nil {
		panic(err)
	}

	// private key of the deployer
	privateKey, err := crypto.HexToECDSA(accountPrivateKey)
	if err != nil {
		panic(err)
	}

	// extract public key of the deployer from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// address of the deployer
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// chain id of the network
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}

	// Get Transaction Ops to make a valid Ethereum transaction
	auth, err := getNextTransaction(client, fromAddress, privateKey, chainID)
	if err != nil {
		panic(err)
	}

	// deploy the contract
	address, tx, simpleStorageApi, err := api.DeployApi(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Api contract deployed to %s\n", address.Hex())
	fmt.Printf("Tx: %s\n", tx.Hash().Hex())

	return auth, client, fromAddress, privateKey, chainID, simpleStorageApi

}

func createWallet(client *ethclient.Client, fromAddress common.Address, privateKey *ecdsa.PrivateKey, chainID *big.Int, simpleStorageApi *api.Api, nameWallet string, balance int64) {
	auth, err := getNextTransaction(client, fromAddress, privateKey, chainID)
	if err != nil {
		panic(err)
	}

	//call contract create wallet First
	_, err = simpleStorageApi.SetWallet(auth, nameWallet, big.NewInt(balance))
	if err != nil {
		panic(err)
	}
}

func infoWallet(simpleStorageApi *api.Api, nameWallet string) {
	//call contract methods read wallet info
	getValues, err := simpleStorageApi.GetWallet(&bind.CallOpts{}, nameWallet)
	if err != nil {
		panic(err)
	}
	fmt.Println(getValues.WalletName)
	fmt.Println(getValues.Balance.String())
}

func transferMoney(client *ethclient.Client, fromAddress common.Address, privateKey *ecdsa.PrivateKey, chainID *big.Int, simpleStorageApi *api.Api, nameWalletSender string, nameWalletRecipient string, money int64) {
	auth, err := getNextTransaction(client, fromAddress, privateKey, chainID)
	if err != nil {
		panic(err)
	}
	//transfer money
	_, err = simpleStorageApi.SendMoney(auth, nameWalletSender, nameWalletRecipient, big.NewInt(money))
	if err != nil {
		panic(err)
	}
}
