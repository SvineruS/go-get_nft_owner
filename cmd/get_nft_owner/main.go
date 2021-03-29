package main

import (
	"fmt"
	ERC721 "get_nft_owner/assets/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const InfuraUrl = "https://mainnet.infura.io/v3/01117e8ede8e4f36801a6a838b24f36c"

func getOwner(hexAddress string, tokenId int) (*common.Address, error) {
	client, err := ethclient.Dial(InfuraUrl)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(hexAddress)
	instance, err := ERC721.NewERC721(address, client)
	if err != nil {
		return nil, err
	}

	tokenId_ := big.NewInt(int64(tokenId))

	owner, err := instance.OwnerOf(nil, tokenId_)
	if err != nil {
		return nil, err
	}

	return &owner, nil
}

func main() {
	owner, err := getOwner("0x06012c8cf97bead5deae237070f9587f8e7a266d", 1997323)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Owner is", owner)
}
