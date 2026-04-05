package main

import "fmt"

func main() {
	var a, b int
	var hun, ten, one int

	fmt.Scan(&a, &b)
	hun = b / 100
	ten = (b % 100) / 10
	one = b % 10

	fmt.Println(one * a)
	fmt.Println(ten * a)
	fmt.Println(hun * a)

	fmt.Println(one*a + ten*a*10 + hun*a*100)
}
