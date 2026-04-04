package main

import (
	"fmt"
	"os"
)

const (
	BANK_FILE = "banks.txt"
)

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

func main() {
	bytes, err := os.ReadFile(BANK_FILE)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return
	}
	fmt.Println(string(bytes))
}
