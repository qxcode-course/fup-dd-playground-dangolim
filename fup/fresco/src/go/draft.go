package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func vogal(c byte) bool {
	return strings.ContainsRune("aeiou", rune(c))
}

func main() {
	in := bufio.NewReader(os.Stdin)

	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)

	p := strings.Split(s, " ")

	res := p[0]

	for i := 1; i < len(p); i++ {
		a := res[len(res)-1]
		b := p[i][0]

		if vogal(a) && vogal(b) {
			j := 0

			// remove apenas as vogais repetidas do início
			for j < len(p[i]) && vogal(p[i][j]) && p[i][j] == a {
				j++
			}

			res += p[i][j:]
		} else {
			res += " " + p[i]
		}
	}

	fmt.Println(res)
}