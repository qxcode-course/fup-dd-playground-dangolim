package main
import "fmt"
func main() {
    var x, y, z int
    fmt.Scan(&x, &y, &z)
    if (x > y && x < z) || (x > z && x < y) {
        fmt.Println(x)
    } else  if (y > x  && y < z) || (y > z && y < x) {
        fmt.Println(y)
    } else {
        fmt.Println(z)  
    }
}  