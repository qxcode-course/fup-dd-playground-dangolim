package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    for i := 1; i <= n; i += 2 {
        fmt.Println(i)
    } 
    start := n
    if start%2 != 0 {
        start--
    }
    for i:= start; i >= 0; i -= 2 {
        fmt.Println(i)
    }
}