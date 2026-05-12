package main
import "fmt"
func main() {
    var maisNovo, quantidade int
    fmt.Scan(&maisNovo, &quantidade)
    for i := 0; i < quantidade; i++ {
        fmt.Println(maisNovo + (i * 2))
    }
}
