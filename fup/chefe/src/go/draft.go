package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Lê código do Ultron
	ultron, _ := reader.ReadString('\n')
	ultron = strings.TrimSpace(strings.ToLower(ultron))

	// Lê linha com as pessoas
	linha, _ := reader.ReadString('\n')
	linha = strings.TrimSpace(strings.ToLower(linha))

	pessoas := strings.Split(linha, " ")

	// Guarda letras do Ultron
	letras := make(map[rune]bool)

	for _, c := range ultron {
		letras[c] = true
	}

	var resposta []string

	for _, pessoa := range pessoas {

		correspondencias := 0

		for _, c := range pessoa {
			if letras[c] {
				correspondencias++
			}
		}

		percentual := float64(correspondencias) / float64(len(pessoa))

		if percentual == 1.0 {
			resposta = append(resposta, "chefe")
		} else if percentual > 0.5 {
			resposta = append(resposta, "ultron")
		} else {
			resposta = append(resposta, "pessoa")
		}
	}

	fmt.Println(strings.Join(resposta, " "))
}