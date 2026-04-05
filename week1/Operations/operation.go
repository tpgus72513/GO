// 첫째 줄에 A+B, 둘째 줄에 A-B, 셋째 줄에 A*B, 넷째 줄에 A/B, 다섯째 줄에 A%B를 출력한다.
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
	fmt.Fprintln(os.Stdout, a-b)
	fmt.Fprintln(os.Stdout, a*b)
	//var c float64 = float64(a) / float64(b)
	//fmt.Fprintln(os.Stdout, c)

	fmt.Fprintln(os.Stdout, a/b)

	//var d float64 = float64(a)%float64(b)
	fmt.Fprintln(os.Stdout, a%b)

}
