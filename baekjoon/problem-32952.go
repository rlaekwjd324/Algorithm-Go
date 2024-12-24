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

	var r, k, m int
	fmt.Fscanln(reader, &r, &k, &m)

	y := 0
	for {
		y += k
		if y > m {
			fmt.Fprintln(writer, r)
			return
		}
		r = r / 2
		if r < 1 {
			fmt.Fprintln(writer, r)
			return
		}
	}
}
