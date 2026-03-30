package main
import "fmt"
func main() {
    var wifi, login, adm  int
    fmt.Scan(&wifi, &login, &adm)
    if wifi== 0 {
        fmt.Print("you must connect to wifi\n")   
        return
}  
if login==0 {
    fmt.Print("you need to login first\n")
     return
}  
if adm == 0 {
    fmt.Print("you must to login as admin\n")  
    return  
} 
fmt.Println("done")
}
