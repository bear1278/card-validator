package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	banks, err := loadBankData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Banks: %#v\n", banks)
}

func loadBankData() ([]Bank, error) {
	var Banks []Bank
	file, err := os.Open(BANK_FILE)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error loading bank file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, errors.New("error loading bank file")
		}
		bankTo, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, errors.New("error loading bank file")
		}
		bankFrom, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, errors.New("error loading bank file")
		}
		Banks = append(Banks, Bank{parts[0], bankTo, bankFrom})
	}
	return Banks, nil
}
