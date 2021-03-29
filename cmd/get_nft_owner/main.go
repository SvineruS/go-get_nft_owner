package main

import (
	"fmt"
	"get_nft_owner/pkg/get_nft_owner"
	"log"
)

func main() {
	owner, err := get_nft_owner.GetOwner("0x06012c8cf97bead5deae237070f9587f8e7a266d", 1997323)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Owner is", owner)
}
