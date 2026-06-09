package main
import (
	"bufio"
	"fmt"
	"os"
)
func ehVogal(c rune)bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscanln(reader, &n)
    for i := 0; i < n; i++ {
        var texto string
        fmt.Fscanln(reader, &texto)
        maior := ""
        atual := ""
        for _, c := range texto {
            if ehVogal(c) {
                atual += string(c)

                if len(atual) > len(maior) {
                    maior = atual
                }
            } else {
                atual =""
            }
        }

        fmt.Println(maior)
    }
}