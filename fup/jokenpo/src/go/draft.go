package main

import "fmt"

func main() {
	var jog1, jog2 string
	fmt.Scan(&jog1, &jog2)
	r := "R"
	p := "P"
	s := "S"
	    if (jog1 == r && jog2==s) || (jog1==s && jog2 == p) || (jog1== p && jog2== r){
        fmt.Println(jog1) 
    } else if (jog2 == r && jog1 ==s) || (jog2==s && jog1 == p) || (jog2== p && jog1== r){
        fmt.Println(jog2)
    } else {
        fmt.Println("empate")
    }
}
