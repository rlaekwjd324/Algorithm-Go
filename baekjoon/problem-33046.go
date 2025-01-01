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

	var a, b, c, d int
	fmt.Fscan(reader, &a, &b, &c, &d)

	f := (a + b) % 4
	if f == 0 {
		f = 4
	}
	s := (c + d - (4 - f + 1))
	if s <= 0 {
		s += 4
	}
	s = s % 4
	if s == 0 {
		s = 4
	}
	fmt.Fprintln(writer, s)
}
