package cli

import (
	"fmt"
	"log"

	"github.com/atsuyaourt/blockchain/internal/blockchain"
)

func (cli *CLI) createBlockchain(address string) {
	if !blockchain.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address)
	bc.DB.Close()
	fmt.Println("Done!")
}
