package main

import (
	"fmt"
)

func main() {
	var soma int
	_, err := fmt.Scan(&soma)
	if err != nil {
		return
	}

	if soma == 0 {
		fmt.Println("joguem de novo")
		return
	}

	indice := (soma - 1) % 26
	letra := rune(97 + indice)

	fmt.Printf("%c\n", letra)
}