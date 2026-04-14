package main
import "fmt"
func main() {
	var m, a, b int
	fmt.Scan(&m, &a, &b)
	c := m - (a + b)
    maior := a
	if b > maior {
		maior = b
	}
	if c > maior {
		maior = c
	}
	fmt.Println(maior)
}