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
		var x, w int
		fmt.Fscan(reader, &x, &w)

		c := 0
		for {
			if x >= w {
				break
			}
			c++
			x *= 2
		}
		
		fmt.Fprintf(writer, "%v\n", c)
	}
}
