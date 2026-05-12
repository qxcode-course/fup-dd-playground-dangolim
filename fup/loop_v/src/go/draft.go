package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    i := a
    first := true
    for {
        if i >= b {
            break
        }
        if i%2 == 0 {
            i++
            continue
        }
        if !first {
            first = false

            i++
        }
        fmt.Println(" ]")
    }
}