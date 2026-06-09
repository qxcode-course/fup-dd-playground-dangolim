package main

import (
	"fmt"
)

func main() {
	var vel, tempoMin, consumo float64
	_, err := fmt.Scan(&vel, &tempoMin, &consumo)
	if err != nil {
		return
	}

	tempoHoras := tempoMin / 60.0
	distancia := vel * tempoHoras
	desempenho := distancia / consumo

	fmt.Printf("%.2f\n", desempenho)
}