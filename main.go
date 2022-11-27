package main

import (
	"github.com/kyu-coder/nomadcoin/cli"
	"github.com/kyu-coder/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
