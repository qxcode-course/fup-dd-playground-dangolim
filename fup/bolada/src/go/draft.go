package main

import (
	"fmt"
)

func main() {
	var tipo string
	var f int
	_, err := fmt.Scan(&tipo, &f)
	if err != nil {
		return
	}

	var t int
	if tipo == "b" {
		t = 20
	} else if tipo == "c" {
		t = 18
	}

	p := ((f * t) - 80) / 10

	if p < 150 {
		fmt.Println("Fraco, nem passou")
	} else if p < 180 {
		fmt.Println("Perfeito")
	} else if p <= 210 {
		fmt.Println("Satisfeito")
	} else {
		fmt.Println("Muito forte, bola fora")
	}
}