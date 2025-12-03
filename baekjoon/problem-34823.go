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

	var y, c, p int
	fmt.Fscan(reader, &y, &c, &p)
	
	c = c/2
	min := y
	if c < min {
		min = c
	}
	if p < min {
		min = p
	}
	
	fmt.Fprintf(writer, "%v\n", min)
}
