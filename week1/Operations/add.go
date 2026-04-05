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
	fmt.Fprintln(os.Stdout, a+b)
}
