package main

import (
	"fmt"
	ERC721 "get_nft_owner/assets/sol"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func query(hexAddress string, tokenId int) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/01117e8ede8e4f36801a6a838b24f36c")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(hexAddress)
	instance, err := ERC721.NewERC721(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded", instance)

	tokenId_ := big.NewInt(int64(tokenId))

	owner, err := instance.OwnerOf(nil, tokenId_)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(owner)
}

func main() {
	query("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb", 6615)
}
