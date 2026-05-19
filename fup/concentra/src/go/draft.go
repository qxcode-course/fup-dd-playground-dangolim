package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    conte := 0 
    for i := a-1; i < b; i++ {
        fmt.Print(i + 1, " ")
        fmt.Print(b - conte, " ")
        conte++
    }
    fmt.Println("]")
    }
     
    

