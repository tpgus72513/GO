package main

import (
	"fmt"
)

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	if m >= 45 {
		fmt.Println(h, m-45)
	} else {
		h -= 1
		m = 60 + (m - 45)
		if h < 0 {
			h += 24
		}
		fmt.Println(h, m)
	}
}

