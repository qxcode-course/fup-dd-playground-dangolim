package main

import (
	"fmt"
)

func main() {
	var h, m, d int
	var sentido string

	_, err := fmt.Scan(&h, &m, &sentido, &d)
	if err != nil {
		return
	}

	minutosIniciais := h*60 + m
	minutosCaminhados := d * 10

	var minutosFinais int
	if sentido == "H" {
		minutosFinais = (minutosIniciais + minutosCaminhados) % 720
	} else if sentido == "A" {
		minutosFinais = (minutosIniciais - minutosCaminhados) % 720
		minutosFinais = (minutosFinais + 720000) % 720
	}

	hFinal := minutosFinais / 60
	mFinal := minutosFinais % 60

	fmt.Printf("%02d %02d\n", hFinal, mFinal)
}