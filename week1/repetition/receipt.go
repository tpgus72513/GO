package main

import (
	"fmt"
)

func main() {
	var X, N int
	fmt.Scan(&X, &N)
	

	totalPrice := 0

	for i := 0; i < N; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		totalPrice += a * b
	}

	if totalPrice == X {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
