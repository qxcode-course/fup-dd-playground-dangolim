package main
import "fmt"
func main() {
    var n1 int
    var n2 int
    var x string
    var resultado int 
    fmt.Scan(&n1, &n2, &x)
    if x == "+" {
       resultado= n1 + n2
    } else if  x == "-"  {
        resultado = n1 - n2
    } else if x == "*" {
        resultado = n1 * n2 
    } else if x == "/" {
        resultado = n1/n2
    }
    fmt.Println(resultado)
}
