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

	var n, c int
	fmt.Fscan(reader, &n, &c)
	
	min := 10201
	
	for i := 0; i < n; i++ {
		var p, d, v int
		fmt.Fscan(reader, &p, &d, &v)
		
		t := p+d+v*c
		if min > t {
			min = t
		}
	}

	fmt.Fprintf(writer, "%v\n", min)
}
