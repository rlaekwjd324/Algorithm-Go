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

    var q int
    fmt.Fscan(reader, &q)
	
	for i := 0; i < q; i++ {
		var l, li, r, c, v int
		fmt.Fscan(reader, &l, &li, &r, &c)

		if l > 2 {
			fmt.Fprintf(writer, "%v\n", "NO")
			continue
		}
		if l == 1 {
			v = 100
		}
		if l == 2 {
			v = 90
		}
		if c < 50 {
			fmt.Fprintf(writer, "%v\n", "NO")
			continue
		}
		if (li * 3 < v && r * 3 < v) || (li <= 21 || r <= 21){
			fmt.Fprintf(writer, "%v\n", "YES")
			continue
		}
		fmt.Fprintf(writer, "%v\n", "NO")
	}
}
