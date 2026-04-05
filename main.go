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
	//banks, err := loadBankData(BANK_FILE)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	var binStr string = ""
	fmt.Scan(&binStr)
	fmt.Println(validateLuhn(binStr))
	//bin := extractBIN(binStr)
	//fmt.Println("Bin:", bin)
	//fmt.Println(identifyBank(bin, banks))

}

func loadBankData(path string) ([]Bank, error) {
	var Banks []Bank
	file, err := os.Open(path)
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
func extractBIN(cardNumber string) int {
	binStr := cardNumber[:6]
	bin, err := strconv.Atoi(binStr)
	if err != nil {
		return 0
	}
	return bin
}

func identifyBank(bin int, banks []Bank) string {
	for _, bank := range banks {
		if bank.BinFrom <= bin && bank.BinTo >= bin {
			return bank.Name
		}
	}
	return "Неизвестный банк"
}

func validateLuhn(cardNumber string) bool {
	var sum int = 0
	if len(cardNumber) != 16 {
		return false
	}
	for i := 16; i > 0; i-- {
		value := int(cardNumber[i-1] - '0')
		if i%2 != 0 {
			value *= 2
		}
		if value > 9 {
			value -= 9
		}
		sum += value
	}
	return sum%10 == 0
}
