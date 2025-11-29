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

	var n, x int
	fmt.Fscan(reader, &n, &x)
	
	sum := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		sum += a
	}

	c := 0
	for {
		if sum/(c+n) >= x {
			fmt.Fprintf(writer, "%v\n", c)
			return
		}
		c++
		sum += 100
	}
	
}
