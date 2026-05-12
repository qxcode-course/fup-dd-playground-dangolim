package main
import "fmt"
func main() {
    var a, b int 
    fmt.Scan(&a, &b)
    for i := a; i < b; i++ {
        fmt.Println(i);
    }
}
