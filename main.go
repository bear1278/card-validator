package main

import (
	"bufio"
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
	file, err := os.Open(BANK_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
