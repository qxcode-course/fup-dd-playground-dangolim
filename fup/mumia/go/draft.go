package main
import "fmt"
func main() {
  var y string
  var x int
  fmt.Scan(&y, &x)
    if x<12 {
      fmt.Println(y,  "eh crianca")
    } else if x<18 && x>=12 {
      fmt.Println(y, "eh jovem")   
    } else  if x<65 && x>=18{
      fmt.Println(y, "eh adulto")   
    } else if x<1000 && x>=65 {
      fmt.Println(y, "eh idoso")   
    } else {
      fmt.Println(y, "eh mumia")   
    }
  }

