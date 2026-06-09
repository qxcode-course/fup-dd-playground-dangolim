package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func vogal(c rune) bool {
	c = unicode.ToLower(c)

	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func gagueja(p string) string {

	r := []rune(p)

	for i := 0; i < len(r)-1; i++ {

		if vogal(r[i]) && !vogal(r[i+1]) {

			silaba := string(r[:i+1])

			return silaba + silaba + p
		}
	}

	return p
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	texto, _ := reader.ReadString('\n')

	palavra := ""

	for _, c := range texto {

		if unicode.IsLetter(c) {
			palavra += string(c)
		} else {

			if palavra != "" {
				fmt.Print(gagueja(palavra))
				palavra = ""
			}

			fmt.Printf("%c", c)
		}
	}

	if palavra != "" {
		fmt.Print(gagueja(palavra))
	}
}