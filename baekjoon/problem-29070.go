package main

import (
    "bufio"
    "fmt"
    "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func ceilDiv(x, y int) int {
	return (x + y - 1) / y
}

func main() {
    defer writer.Flush()

	var a, b, h, w int
	fmt.Fscan(reader, &a, &b, &h, &w)
	
	v := ceilDiv(h, a)*ceilDiv(w, b)

	fmt.Fprintf(writer, "%v\n", v)
}
