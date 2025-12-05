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

	var t int
	fmt.Fscan(reader, &t)
	
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		c := 0
		for {
			c += n%2
			n = n/2+n%2
			if n < 2 {
				break
			}
		}
		fmt.Fprintf(writer, "%v\n", c)
	}
}
