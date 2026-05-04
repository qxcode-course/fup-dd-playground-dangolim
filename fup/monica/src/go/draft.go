package main
import "fmt"
func main() {
    var m, a, b int
    fmt.Scan(&m)
    fmt.Scan(&a)
    fmt.Scan(&b)
    c:= m - a - b
    maior:= a
    if b > maior {
        maior = b
    } 
    if c > maior {
        maior = c
    }
    fmt.Println(maior)
}
