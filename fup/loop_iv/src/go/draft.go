package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    step := 1
    if a > b {
        step = -1
    }
    first := true
    for i := a; i != b; i += step {
        if !first {
            fmt.Print(" ")
        }
        fmt.Print(i)
        first= false
    }
    fmt.Println(" ]")
}