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

	c := 0
	for {
		if n == 0 {
			break
		}
		c += (n%2)
		n = n / 2
	}
	
	fmt.Fprintf(writer, "%v\n", c)
}
