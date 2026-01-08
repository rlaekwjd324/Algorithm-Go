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
	
	min, max, sum := 2000000000.0, 0.0, 0.0

	for i := 0; i < n; i++ {
		var a, b float64
		fmt.Fscan(reader, &a, &b)
		
		if max < a/b {
			max = a/b
		}
		if min > a/b {
			min = a/b
		}
		sum += a/b
	}

	fmt.Fprintf(writer, "%.11f %.11f %.11f\n", min, max, sum)
}
