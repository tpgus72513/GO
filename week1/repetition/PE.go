package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < (n / 4); i++ {
		fmt.Printf("long ")
	}
	fmt.Printf("int")
}
