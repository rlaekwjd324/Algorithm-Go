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

	var n, a, b, c int
	fmt.Fscan(reader, &n, &a, &b, &c)

	d := 0
	if n == 1 {
		d = 0
	} else {
		d = min(a, b) + (n - 2) * min(min(a, b), c)
	}

	h := d/100
	d = d%100
	fmt.Fprintf(writer, "%v %v\n", h, d)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
