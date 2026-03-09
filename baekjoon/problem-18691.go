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
	
	for i := 0; i < n; i++ {
		var g, c, e int
		fmt.Fscan(reader, &g, &c, &e)
		
		v := (e-c)*(2*g-1)
		if c >= e {
			v = 0
		}
		fmt.Fprintf(writer, "%v\n", v)
	}
}
