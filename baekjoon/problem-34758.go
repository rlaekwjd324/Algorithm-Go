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

	var x, y, n int
	fmt.Fscan(reader, &x, &y, &n)
	
	for i := 0; i < n; i++ {
		var tx, ty int
		fmt.Fscan(reader, &tx, &ty)
		
		if tx == x || ty == y {
			fmt.Fprintf(writer, "%v\n", 0)
			continue
		}
		fmt.Fprintf(writer, "%v\n", 1)
	}
}
