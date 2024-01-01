package cli

import (
	"fmt"

	"github.com/atsuyaourt/blockchain/internal/blockchain"
)

func (cli *CLI) reindexUTXO() {
	bc := blockchain.NewBlockchain()
	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
