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

	pre := int64(-1000000000000000000)
	for i := 0; i < n; i++ {
		var a int64
		fmt.Fscan(reader, &a)

		if pre >= a {
			fmt.Fprintf(writer, "%v\n", 0)
			return
		}
		pre = a
	}
		
	fmt.Fprintf(writer, "%v\n", 1)
}
