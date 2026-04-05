package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	for {
		binStr := getUserInput()
		if len(binStr) == 0 {
			break
		}
		if !validateInput(binStr) {
			fmt.Println("Invalid input")
			continue
		}
		fmt.Println(validateLuhn(binStr))
	}
}

func getUserInput() string {
	fmt.Print("Введите номер карты (или Enter для выхода):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func validateInput(cardNumber string) bool {
	if len(cardNumber) < 13 || len(cardNumber) > 19 {
		return false
	}
	for _, char := range cardNumber {
		if !unicode.IsNumber(char) {
			return false
		}
	}
	return true
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
	for i := 1; i <= len(cardNumber); i++ {
		value := int(cardNumber[len(cardNumber)-i] - '0')
		if i%2 == 0 {
			value *= 2
		}
		if value > 9 {
			value -= 9
		}
		sum += value
	}
	return sum%10 == 0
}
