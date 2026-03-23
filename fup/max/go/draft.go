package main
import (
    "fmt"
)
func main() {
    var n1, n2  float64
    fmt.Scan(&n1, &n2)
    if n1 > n2 {
        fmt.Println(n1)
    } else if n2 > n1 {
        fmt.Println(n2)
    } else {
        fmt.Println(n1)
    }
}
