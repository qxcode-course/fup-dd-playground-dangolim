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

	letraStr, _ := reader.ReadString('\n')
	letraStr = strings.TrimSpace(letraStr)

	letra := unicode.ToLower(rune(letraStr[0]))

	contador := 0

	for _, c := range frase {
		if unicode.ToLower(c) == letra {
			contador++
		}
	}

	fmt.Println(contador)
}