package main

import "fmt"

func main() {
	name := "Kyu ! ! ! ! ! ! ! Is my name"
	for _, letter := range name {
		fmt.Printf("%x\n", letter)
	}
}
