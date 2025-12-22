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

	m := 0
	a := ""
	for i := 0; i < n; i++ {
		var s string
		var y int
		fmt.Fscan(reader, &s, &y)

		if m < y {
			m = y
			a = s
		}
	}
		
	fmt.Fprintf(writer, "%v\n", a)
}
