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

	var n, k int
	fmt.Fscan(reader, &n, &k)
	
	c := 0
	max := 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		
		c += a
		c -= b
		if c < 0 {
			c = 0
		}
		if max < c-k {
			max = c-k
		}
	}

	fmt.Fprintf(writer, "%v\n", max)
}
