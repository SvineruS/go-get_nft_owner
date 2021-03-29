solc --abi --bin ERC721.sol -o build --overwrite
abigen --bin=./build/ERC721.bin --abi=./build/ERC721.abi --pkg=ERC721 --out=ERC721.go
