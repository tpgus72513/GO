package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(reader, &a, &b)
	var c float64
	c = float64(a) / float64(b)
	fmt.Fprintln(os.Stdout, c)
}
