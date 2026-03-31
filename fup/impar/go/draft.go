package main
import "fmt"
func main() {
    var p, d1, d2 int
    fmt.Scan(&p, &d1, &d2)
    if p == 0 {
        if (d1 + d2)% 2 == 0 {
             fmt.Println("0") 
        } else {
        fmt.Println("1")
       }
    } else {
           if (d1 + d2)% 2 == 0 {
             fmt.Println("1") 
        } else {
            fmt.Println("0") } }
}
