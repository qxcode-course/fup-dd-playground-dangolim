package main
import "fmt"
import "math"
func main() {
    var c string
    var x float64
    fmt.Scan(&c, &x)
    if c == "r" {
        fmt.Println(math.Round(x)) 

    } else if c == "f" {
        fmt.Println(math.Floor(x)) 

    } else {
        fmt.Println(math.Ceil(x)) 

    }
}
