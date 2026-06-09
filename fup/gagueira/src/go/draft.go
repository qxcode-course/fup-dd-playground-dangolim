package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	frase := scanner.Text()

	palavras := strings.Split(frase, " ")

	for i, palavra := range palavras {
		fmt.Print(palavra, " ", palavra)

		if i != len(palavras)-1 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
}