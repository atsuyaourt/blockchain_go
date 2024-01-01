package cli

import (
	"fmt"

	"github.com/atsuyaourt/blockchain/internal/blockchain"
)

func (cli *CLI) createWallet() {
	wallets, _ := blockchain.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
