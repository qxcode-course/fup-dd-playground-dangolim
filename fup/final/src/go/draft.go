package main
import "fmt"
func main() {
    var n1, n2, n3  int
    fmt.Scan(&n1, &n2, &n3 )
    media := (n1+n2)/2 
    if media >= 7 {
        fmt.Println("aprovado")
    } else if media < 4 {
        fmt.Println("reprovado")
    } else {
        novamedia := (media+n3)/2 
      if novamedia >=5 {
        fmt.Println("aprovado na final")
    } else {
        fmt.Println("reprovado na final")
    }
}
}