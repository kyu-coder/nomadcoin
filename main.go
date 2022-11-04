package main

import (
	"github.com/kyu-coder/nomadcoin/blockchain"
	"github.com/kyu-coder/nomadcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
