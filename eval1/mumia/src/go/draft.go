package main
import "fmt"
func main() {
    var nome string
    var x int
    fmt.Scan(&nome, &x)
    if x<12 {
        fmt.Println(nome,"eh crianca")
    } else if x>=12 && x<18 {
        fmt.Println(nome,"eh jovem")
    } else if x>=18 && x<65 {
        fmt.Println(nome,"eh adulto")
    } else if x>=65 && x<1000 {
        fmt.Println(nome,"eh idoso")
    } else {
        fmt.Println(nome,"eh mumia")
    }
}
