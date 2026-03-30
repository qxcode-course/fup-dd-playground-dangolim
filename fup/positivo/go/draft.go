package main
import "fmt"
func main() {
    var n1 int
    fmt.Scan(&n1)
    if n1 >= 0 {
        fmt.Println("SIM")
    } else if n1 < 0 {
        fmt.Println(" ")
    }
    
}
