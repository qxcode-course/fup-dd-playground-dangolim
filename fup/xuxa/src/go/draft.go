package main

import (
	"bufio"
	"fmt"
	"os"
)
func inverterFrase(s string) string {
    runes:= []rune(s)
    for i, j := 0, len(runes)-1; i<j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
    func main(){
        scanner := bufio.NewScanner(os.Stdin)
        var frase string 
        scanner.Scan()
        frase = scanner.Text()
        frase = inverterFrase(frase)
        fmt.Println(frase)
    }
