package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"bridge/bridge"
	"bridge/ierc20"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	bridgeABI       abi.ABI
	settings        Settings
	zeroAddress     common.Address    = common.HexToAddress("0x0000000000000000000000000000000000000000")
	sequenceNumbers map[uint64]uint64 = make(map[uint64]uint64)
)

type eventDecoded struct {
	Destination        common.Address
	Amount             *big.Int
	Token              common.Address
	OriginChainID      uint64
	DestinationChainID uint64
}

func main() {
	loadSequenceNumbers()
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
		settings.ChainIDToABI[chainID], err = bridge.NewBridge(address, settings.ChainIDToClients[chainID])
		if err != nil {
			log.Fatalln(err)
		}
		//listen
		go monitorBridge(chainID, settings.ChainIDToABI[chainID])

	}
	ticker := time.NewTicker(1 * time.Hour)
	counter := 0
	for {
		select {
		case _ = <-ticker.C:
			counter++
			log.Println(fmt.Sprintf("Uptime: %d hours", counter))
		}
	}
}
func monitorBridge(chainId uint64, scopedBridge *bridge.Bridge) {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			request, err := scopedBridge.CheckForValidRequest(&bind.CallOpts{Pending: false}, big.NewInt(int64(sequenceNumbers[chainId])))
			if err != nil {
				log.Println(err)
			} else if request.Amount.Cmp(big.NewInt(0)) == 1 {
				eventData := eventDecoded{}
				eventData.Amount = request.Amount
				eventData.Destination = request.Recipient
				//eventData.Token = request.Token
				eventData.OriginChainID = request.OriginChainID
				eventData.DestinationChainID = request.DestinationChainID
				FinishBridge(eventData)
			}
		}
	}

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
		} else {
			log.Println("TX sent: ", tx.Hash())
			sequenceNumbers[eventData.OriginChainID]++
			saveSequenceNumbers()
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
			log.Println("TX sent: ", tx.Hash())
			sequenceNumbers[eventData.OriginChainID]++
			saveSequenceNumbers()
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
func saveSequenceNumbers() {
	jsonString, _ := json.Marshal(sequenceNumbers)
	os.WriteFile("sequenceNumbers.json", jsonString, os.ModePerm)
}
func loadSequenceNumbers() {
	file, _ := os.ReadFile("sequenceNumbers.json")
	json.Unmarshal(file, &sequenceNumbers)
}
