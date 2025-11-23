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

	var n, m float64
	fmt.Fscan(reader, &n, &m)

	if m/n*100 >= 81.0 {
		fmt.Fprintf(writer, "%v\n", "yaho")
		return
	}
	
	fmt.Fprintf(writer, "%v\n", "no")
}
