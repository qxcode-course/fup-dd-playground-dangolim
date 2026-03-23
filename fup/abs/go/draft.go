package main
import  (
    "fmt"
    "math"
)
func main() {
    var n1, n2, valorabs  float64
    fmt.Scan(&n1, &n2, &valorabs)
    valorabs= math.Abs(n1-n2)
    fmt.Println(valorabs)
}
