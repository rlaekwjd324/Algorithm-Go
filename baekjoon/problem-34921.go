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

	var a, t int
	fmt.Fscan(reader, &a, &t)
	
	p := 10 + 2 * (25 - a + t)
	
	if p < 0 {
		p = 0
	}
	
	fmt.Fprintf(writer, "%v\n", p)
}
