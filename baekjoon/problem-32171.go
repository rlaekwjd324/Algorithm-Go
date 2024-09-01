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

	minX := 10
	minY := 10
	maxX := -10
	maxY := -10
	for i := 0; i < n; i++ {
		var a, b, c, d int
		fmt.Fscan(reader, &a, &b, &c, &d)

		if minX > a {
			minX = a
		}
		if minY > b {
			minY = b
		}
		if maxX < c {
			maxX = c
		}
		if maxY < d {
			maxY = d
		}

		sum := ((maxX-minX)*2 + (maxY-minY)*2)

		fmt.Fprintln(writer, sum)
	}
}
