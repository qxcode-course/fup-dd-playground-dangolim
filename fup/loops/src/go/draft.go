package main

import "fmt"

func main() {
	var ang int

	fmt.Scan(&ang)

	ang = ((ang % 360) + 360) % 360

	fmt.Println(ang)
}