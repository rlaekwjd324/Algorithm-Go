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
	fmt.Fscanln(reader, &n)

	c := 0
	max := 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)

		c -= a
		c += b
		if max < c {
			max = c
		}
	}

	fmt.Fprintln(writer, max)
}
