package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	sum := 1
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)

		v := b - a
		if a*b < 0 {
			if b < 0 {
				v++
			} else {
				v--
			}
		}
		sum += v
	}
	if sum <= 0 {
		sum--
	}
	fmt.Fprintf(writer, "%v\n", sum)
}
