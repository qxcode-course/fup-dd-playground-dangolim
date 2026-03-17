package main
import "fmt"
func main() {
    var valor float64
    var celsius, fahrenheit float64
    fmt.Scan(&celsius, &fahrenheit, &valor)
    fahrenheit= (1.8 * celsius) + 32
    valor= fahrenheit
    fmt.Printf("%.6f\n", valor)
}
