package main

import "fmt"

func main() {
	var x int = 5
	y := x + 2
	if y > 10 {
		fmt.Println("y is greater than 10")
	} else if y == 10 {
		fmt.Println("y is equal to 10")
	} else {
		fmt.Println("y is less than 10")
	}
}
