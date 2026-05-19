package main
import "fmt"
func main() {
    var c, m, passageiros int
    fmt.Scan(&c)
    for {
        qtd, err := fmt.Scan(&m)
        if err != nil || qtd == 0 {
            break
        }
        passageiros += m
        if passageiros == 0 {
        fmt.Println("vazio")
    } else if passageiros < c {
        fmt.Println("ainda cabe")
    } else if passageiros >= c && passageiros< 2*c {
        fmt.Println("lotado")
    } else if passageiros >= 2*c{
        fmt.Println("hora de partir")
        break
    }
    }
}