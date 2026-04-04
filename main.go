package main

import "fmt"

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

func main() {
	fmt.Println(Bank{})
	fmt.Println(Bank{"Lunar Bank", 40000, 49999})
}
