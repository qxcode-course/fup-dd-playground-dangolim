package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var ultron, pessoa string

		fmt.Fscan(in, &ultron)
		fmt.Fscan(in, &pessoa)

		ultron = strings.ToLower(ultron)
		pessoa = strings.ToLower(pessoa)

		// Guarda as letras do código do Ultron
		letras := make(map[rune]bool)

		for _, c := range ultron {
			letras[c] = true
		}

		// Conta quantas letras da pessoa pertencem ao código do Ultron
		correspondencias := 0

		for _, c := range pessoa {
			if letras[c] {
				correspondencias++
			}
		}

		percentual := float64(correspondencias) / float64(len(pessoa))

		if percentual == 1.0 {
			fmt.Println("chefe")
		} else if percentual > 0.5 {
			fmt.Println("ultron")
		} else {
			fmt.Println("pessoa")
		}
	}
}