package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"bridge/bridge"
	"bridge/ierc20"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	bridgeABI   abi.ABI
	settings    Settings
	zeroAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

type eventDecoded struct {
	Destination        common.Address
	Amount             *big.Int
	Token              common.Address
	OriginChainID      uint64
	DestinationChainID uint64
}

func main() {
	contractABI, err := abi.JSON(strings.NewReader(string(bridge.BridgeABI)))
	if err != nil {
		log.Fatal(err)
	}
	bridgeABI = contractABI
	settings = ReadConfig("config.yaml")

	publicKey := settings.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	settings.Address = crypto.PubkeyToAddress(*publicKeyECDSA)

	for address, chainID := range settings.AddressToChainID {
		//listen
		go subscribeToEventLogs(settings.ChainIDToClients[chainID], address)
	}
	ticker := time.NewTicker(60 * time.Hour)
	counter := 0
	for {
		select {
		case _ = <-ticker.C:
			counter++
			log.Println(fmt.Sprintf("Uptime: %d hours", counter))
		}
	}
}
func subscribeToEventLogs(client *ethclient.Client, bridgeAddress common.Address) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{bridgeAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			processLog(vLog)
		}
	}
}
func processLog(vLog types.Log) {
	eventData := eventDecoded{}
	eventDataMapInterface := make(map[string]interface{})
	err := bridgeABI.UnpackIntoMap(eventDataMapInterface, "OrderPlaced", vLog.Data)
	if err != nil {
		fmt.Println("Failed to unpack:", err)
	}
	eventData.Amount = eventDataMapInterface["amount"].(*big.Int)
	eventData.Destination = eventDataMapInterface["recipient"].(common.Address)
	eventData.Token = eventDataMapInterface["token"].(common.Address)
	eventData.OriginChainID = eventDataMapInterface["originChainID"].(*big.Int).Uint64()
	eventData.DestinationChainID = eventDataMapInterface["destinationChainID"].(*big.Int).Uint64()
	FinishBridge(eventData)
}
func FinishBridge(eventData eventDecoded) {
	client := settings.ChainIDToClients[eventData.DestinationChainID]
	if eventData.Token.Cmp(zeroAddress) == 0 {
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Println("Error retrieving gas price")
			return
		}
		nonce, err := client.PendingNonceAt(context.Background(), settings.Address)
		if err != nil {
			log.Println(err)
		}
		tx := types.NewTransaction(nonce, eventData.Destination, eventData.Amount, 21000, gasPrice, []byte{})

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), settings.PrivateKey)
		if err != nil {
			log.Fatal(err)
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		token, err := ierc20.NewIERC20(eventData.Token, client)
		if err != nil {
			log.Fatalf("Failed to instantiate a Token contract: %v", err)
		}
		tx, err := token.Transfer(generateAuth(client, eventData.OriginChainID), eventData.Destination, eventData.Amount)
		if err != nil {
			log.Println("Error transferring tokens: ", err.Error())
		} else {
			log.Println("TX hash: ", tx.Hash())
		}

	}

}

func generateAuth(client *ethclient.Client, chainId uint64) *bind.TransactOpts {
	nonceCurrent, nErr := client.PendingNonceAt(context.Background(), settings.Address)
	if nErr != nil {
		log.Println(nErr)
	} else {
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Println("Error retrieving current block number")
		} else {
			auth, _ := bind.NewKeyedTransactorWithChainID(settings.PrivateKey, big.NewInt(int64(chainId)))

			auth.Nonce = big.NewInt(int64(nonceCurrent))
			auth.Value = big.NewInt(0)      // in wei
			auth.GasLimit = uint64(1000000) // in units
			auth.GasPrice = gasPrice
			return auth
		}
	}
	return nil
}
