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

	var w, l, h float64
	fmt.Fscan(reader, &w, &l, &h)
	
	minWl := w
	maxWl := w
	if w > l {
		minWl = l
	} else {
		maxWl = l
	}

	if minWl/h >= 2 && maxWl/minWl <= 2 {
		fmt.Fprintf(writer, "%v\n", "good")
		return
	}
	fmt.Fprintf(writer, "%v\n", "bad")
}
