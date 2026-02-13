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
	
	c := 0
	f := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		
		c+= a
		if a == 0 {
			f++
		}
	}
	
	if float64(f) >= float64(n)/2.0 {
		fmt.Fprintf(writer, "%v\n", "INVALID")
		return
	}
	if c > 0 {
		fmt.Fprintf(writer, "%v\n", "APPROVED")
		return
	}
	fmt.Fprintf(writer, "%v\n", "REJECTED")
}
