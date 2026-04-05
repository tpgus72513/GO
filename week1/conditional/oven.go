package main

import (
	"fmt"
)

func main() {
	var h, m, time int
	fmt.Scan(&h, &m)
	fmt.Scan(&time)

	m += time

	if m > 59 {
		h = h + (m / 60)
		m %= 60
	}
	if h > 23 {
		h %= 24
	}

	fmt.Println(h, m)
}
