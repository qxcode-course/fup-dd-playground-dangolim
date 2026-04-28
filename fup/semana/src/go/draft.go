package main
import "fmt"
func main() {
   var x, hora int
   fmt.Scan(&x, &hora)
    if x == 1 {
        fmt.Println("NAO")
    } else if x >=2 && x <=6 {
        if hora >=8 && hora <=11 || hora >=14 && hora <= 17 {
            fmt.Println("SIM")
        } else {
            fmt.Println("NAO")
        }
    } else if x == 7 {
        if hora >=8 && hora <=11 {
                fmt.Println("SIM")
            } else {
                fmt.Println("NAO")
            }
        }
    }