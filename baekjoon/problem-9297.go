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
		var n, d int
		fmt.Fscan(reader, &n, &d)
		
		l := n/d
		n = n%d
		
		if n == 0 {
			fmt.Fprintf(writer, "Case %v: %v\n", i+1, l)
			continue
		}
		if l == 0 {
			fmt.Fprintf(writer, "Case %v: %v/%v\n", i+1, n, d)
			continue
		}
		fmt.Fprintf(writer, "Case %v: %v %v/%v\n", i+1, l, n, d)
	}
}
