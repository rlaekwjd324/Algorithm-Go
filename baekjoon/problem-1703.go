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

	for {
		var n int
		fmt.Fscan(reader, &n)

		if n == 0 {
			return
		}
		c := 1
		for i := 0; i < n; i++ {
			var a, b int
			fmt.Fscan(reader, &a, &b)
			
			c = c*a-b
		}
		fmt.Fprintf(writer, "%v\n", c)
	}
}
