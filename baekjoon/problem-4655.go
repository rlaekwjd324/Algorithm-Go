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
		var n float64
		fmt.Fscan(reader, &n)

		if n == 0.0 {
			return
		}

		c := 2.0
		sum := 0.0
		for {
			sum += (1.0/c)

			if sum >= n {
				break
			}
			c++
		}
		
		fmt.Fprintf(writer, "%v card(s)\n", c-1)
	}
}
