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

	texto, _ := reader.ReadString('\n')
	texto = strings.TrimSpace(texto)

	comando, _ := reader.ReadString('\n')
	comando = strings.TrimSpace(comando)

	if comando == "M" {
		fmt.Println(strings.ToUpper(texto))

	} else if comando == "m" {
		fmt.Println(strings.ToLower(texto))

	} else if comando == "i" {

		for _, c := range texto {

			if unicode.IsUpper(c) {
				fmt.Printf("%c", unicode.ToLower(c))
			} else if unicode.IsLower(c) {
				fmt.Printf("%c", unicode.ToUpper(c))
			} else {
				fmt.Printf("%c", c)
			}
		}

		fmt.Println()

	} else if comando == "p" {

		palavras := strings.Split(texto, " ")

		for i := 0; i < len(palavras); i++ {

			p := palavras[i]

			if len(p) == 1 {
				palavras[i] = strings.ToLower(p)
			} else {
				palavras[i] = strings.ToUpper(string(p[0])) + strings.ToLower(p[1:])
			}
		}

		fmt.Println(strings.Join(palavras, " "))
	}
}