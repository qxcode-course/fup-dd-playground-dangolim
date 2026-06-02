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

	frase = strings.TrimRight(frase, "\n")

	for _, ch := range frase {

		if ch == ' ' {
			fmt.Print(" ")
		} else {
			letra := unicode.ToLower(ch)

			if letra == 'a' || letra == 'e' || letra == 'i' || letra == 'o' || letra == 'u' {
				fmt.Print("v")
			} else {
				fmt.Print("c")
			}
		}
	}

	fmt.Println()
}