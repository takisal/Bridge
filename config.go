package main

import (
	"bridge/bridge"
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v2"
)

type Config struct {
	PrivateKey *ecdsa.PrivateKey `yaml:"private_key"`
	ChainIDs   []uint64          `yaml:"chain_ids"`
	Bridges    []common.Address  `yaml:"bridges"`
	RPCs       []string          `yaml:"rpcs"`
}
type Settings struct {
	Address          common.Address
	PrivateKey       *ecdsa.PrivateKey
	ChainIDToAddress map[uint64]common.Address
	AddressToChainID map[common.Address]uint64
	ChainIDToClients map[uint64]*ethclient.Client
	ChainIDToABI     map[uint64]*bridge.Bridge
}

func ReadConfig(filePath string) Settings {
	c := Config{}
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	settings := Settings{PrivateKey: c.PrivateKey}
	for i := 0; i < len(c.Bridges); i++ {
		settings.AddressToChainID[c.Bridges[i]] = c.ChainIDs[i]
		settings.ChainIDToAddress[c.ChainIDs[i]] = c.Bridges[i]
		client, err := ethclient.Dial(c.RPCs[i])
		if err != nil {
			log.Fatal(err)
		}
		settings.ChainIDToClients[c.ChainIDs[i]] = client
	}
	return settings
}
