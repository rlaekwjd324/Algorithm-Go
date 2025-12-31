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

	var n, p float64
	fmt.Fscan(reader, &n, &p)

	y := n * 10
	m := 0.0
	for i := 1; i <= 10; i++ {
		m += n*(1+float64(i)*p/100.0/10.0)
	}
	
	fmt.Fprintf(writer, "%.4f\n", m-y)
}
