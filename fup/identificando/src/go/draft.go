package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)
	elementos := strings.Fields(frase)
	for i, e := range elementos {
		tipo := "int"
		for _, c := range e {
			if unicode.IsLetter(c) {
				tipo = "str"
				break
			}
		}
		if tipo != "str" {
			if strings.Contains(e, ".") {
				tipo = "float"
			}
		}

		if i > 0 {
			fmt.Print(" ")
		}

		fmt.Print(tipo)
	}

	fmt.Println()
}