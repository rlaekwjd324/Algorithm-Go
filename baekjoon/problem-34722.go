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
	
	cnt := 0
	for i := 0; i < n; i++ {
		var s, c, a, r int
		fmt.Fscan(reader, &s, &c, &a, &r)

		if s >= 1000 || c >= 1600 || a >= 1500 || (r <= 30 && r != -1) {
			cnt++
		}
	}

	fmt.Fprintf(writer, "%v\n", cnt)
}
