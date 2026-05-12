package main
import (
    "fmt"
    "bufio"
    "os"
)

func inverterCase(c rune) rune {
    if c >= 'a' && c <= 'z' {
        return c - ('a' - 'A')
    }
    if c >= 'A' && c <= 'Z' {
        return c + ('a' - 'A')
    }
    return c
}
    func main() {
         scanner := bufio.NewScanner(os.Stdin)
        var frase string 
        scanner.Scan()
        frase = scanner.Text()
        frase = string(inverterCase(rune(frase[0])))
        fmt.Println(frase)
        }      
    

