package main

import (
	"fmt"
)

func main() {
	var h, m, d, mes, ano int
	_, err := fmt.Scan(&h, &m, &d, &mes, &ano)
	if err != nil {
		return
	}

	anoDoisDigitos := ano % 100

	fmt.Printf("%02d:%02d %02d/%02d/%02d\n", h, m, d, mes, anoDoisDigitos)
}