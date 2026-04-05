package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
