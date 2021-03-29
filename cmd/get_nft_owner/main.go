package main

import (
	"fmt"
	ERC721 "get_nft_owner/assets/sol"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func query(hexAddress string, tokenId int) {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(hexAddress)
	instance, err := ERC721.NewERC721(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded", instance)
	// todo
}

func main() {
	query("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb", 6615)
}
