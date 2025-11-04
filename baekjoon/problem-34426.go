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
		var n int
		fmt.Fscan(reader, &n)
		
		m := 1000000.0
		idx := -1
		for j := 0; j < n; j++ {
			var a, b float64
			fmt.Fscan(reader, &a, &b)

			v := b/a
			if v < m {
				m = v
				idx = j+1
			}
		}
		
		fmt.Fprintf(writer, "%v\n", idx)
	}

}
