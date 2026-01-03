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

	var n, a, b, s, t int
	fmt.Fscan(reader, &n, &a, &b, &s, &t)

	if a < s && a < t && b > s && b > t {
		fmt.Fprintf(writer, "%v\n", "City")
		return
	}
	if b <= s && b <= t {
		fmt.Fprintf(writer, "%v\n", "Outside")
		return
	}
	if a >= s && a >= t {
		fmt.Fprintf(writer, "%v\n", "Outside")
		return
	}
	fmt.Fprintf(writer, "%v\n", "Full")
}
