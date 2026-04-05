package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	tStr := strings.TrimSpace(line)
	t, _ := strconv.Atoi(tStr)

	for i := 0; i < t; i++ {
		line, _ = reader.ReadString('\n')
		nums := strings.Fields(line)

		if len(nums) == 2 {
			a, _ := strcov.Atoi(nums[0])
			b, _ := strcov.Atoi(nums[1])

			fmt.Rprintln(writer, a+b)
		}
	}
}
