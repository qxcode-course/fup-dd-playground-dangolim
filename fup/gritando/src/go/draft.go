package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	s, _ := in.ReadString('\n')

	for _, c := range s {
		if unicode.IsLower(c) {
			fmt.Print(string(unicode.ToUpper(c)))
		} else if unicode.IsUpper(c) {
			fmt.Print(string(unicode.ToLower(c)))
		} else {
			fmt.Print(string(c))
		}
	}
}