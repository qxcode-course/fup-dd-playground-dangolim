package main

import (
	"fmt"
	"strings"
)

func ehVogal(c byte) bool {
	c = byte(strings.ToLower(string(c))[0])
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func contarVogais(nome string) int {
	contador := 0

	for i := 0; i < len(nome); i++ {
		if ehVogal(nome[i]) {
			contador++
		}
	}

	return contador
}

func main() {
	var nome1, nome2 string
	fmt.Scanln(&nome1)
	fmt.Scanln(&nome2)

	pontos := 0

	// Primeira letra igual
	if strings.ToLower(string(nome1[0])) == strings.ToLower(string(nome2[0])) {
		pontos += 20
	}

	// Mesmo tamanho
	if len(nome1) == len(nome2) {
		pontos += 30
	}

	// Mesma quantidade de vogais
	if contarVogais(nome1) == contarVogais(nome2) {
		pontos += 30
	}

	// Verifica última letra
	final1 := nome1[len(nome1)-1]
	final2 := nome2[len(nome2)-1]

	if (ehVogal(final1) && ehVogal(final2)) || (!ehVogal(final1) && !ehVogal(final2)) {
		pontos += 20
	} else {
		pontos -= 10
	}

	// Não pode ser negativo
	if pontos < 0 {
		pontos = 0
	}

	fmt.Printf("As chances do crush te dar bola sao: %d%%!\n", pontos)
}