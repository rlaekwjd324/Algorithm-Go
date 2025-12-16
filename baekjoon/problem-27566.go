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

	var r, f int
	fmt.Fscan(reader, &r, &f)

	v := 180 * f/r
	v = v % 360

	if v > 90 && v < 270 {
		fmt.Fprintln(writer, "down")
		return
	}
	fmt.Fprintln(writer, "up")
}
