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
		var n, m int
		fmt.Fscan(reader, &n, &m)
		
		v := n*m - (n-1)*(m-1)
		fmt.Fprintf(writer, "%v\n", v)
	}
}
