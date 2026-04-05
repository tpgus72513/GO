package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		fmt.Println(10000 + a*1000)
	} else if a == b || a == c {
		fmt.Println(1000 + a*100)
	} else if b == c {
		fmt.Println(1000 + b*100)
	} else {
		max := a
		if b > max {
			max = b
		}
		if c > max {
			max = c
		}
		fmt.Println(max * 100)
	}
}
