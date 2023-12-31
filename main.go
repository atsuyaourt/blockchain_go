package main

import (
	"github.com/atsuyaourt/blockchain/cmd/cli"
	"github.com/atsuyaourt/blockchain/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.DB.Close()

	cli := cli.NewCLI(bc)
	cli.Run()
}
