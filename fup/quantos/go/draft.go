package main
import "fmt"
func main() {
    var n1, n2, n3 int
    fmt.Scan(&n1, &n2, &n3)
    if n1== n2 && n2== n3  {
        fmt.Println(3)
    } else if n1== n3 || n3== n1 ||  n2== n3 {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
    
}
