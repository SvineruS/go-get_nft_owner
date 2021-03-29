package get_nft_owner

import (
	ERC721 "get_nft_owner/assets/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const InfuraUrl = "https://mainnet.infura.io/v3/01117e8ede8e4f36801a6a838b24f36c"

func GetOwner(hexAddress string, tokenId int) (*common.Address, error) {
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
