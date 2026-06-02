package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)
func main() {
    reader := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscanln(reader, &n)

    for i := 0; i < n; i++ {
        texto, _ := reader.ReadString('\n')
        runas := []rune(texto)
        maiuscula := false
        for _, c := range runas {
            if unicode.IsLetter(c) {
                if unicode.IsUpper(c) {
                    maiuscula = true
                }
                break
            }
        }
        for _, c := range runas {
            if c == ' ' || c == '\n' {
                fmt.Printf("%c", c)
                continue
            }
            if maiuscula {
                fmt.Printf("%c", unicode.ToUpper(c))
            } else {
                fmt.Printf("%c", unicode.ToLower(c))
            }
            maiuscula = !maiuscula
        }
    }
}