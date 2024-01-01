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
	defer bc.DB.Close()

	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
