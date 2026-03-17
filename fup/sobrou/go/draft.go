package main
import "fmt"
func main() {
    var a, b, c, x, y, z, dinheiro float64
    fmt.Scan(&a, &b, &c, &x, &y, &z, &dinheiro)
    sobra:= dinheiro - (a*x + b*y + c*z)
    fmt.Printf("%.2f\n", sobra )
}
