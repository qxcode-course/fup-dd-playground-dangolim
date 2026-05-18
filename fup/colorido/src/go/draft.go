package main
import "fmt"
func main() {
	var N int
	var pe string
	fmt.Scan(&N)
	fmt.Scan(&pe)
	atual := pe
    fmt.Print("[ ")
	for i := 0; i <= 9; i++ {
        
		if i == N {
			continue
		}

		fmt.Printf("%d%s ", i, atual)

		if atual == "d" {
			atual = "e"
		} else {
			atual = "d"
		}
	}
	if N != 10 {
		fmt.Print("ceu ")
	}
	fmt.Println("]")
}