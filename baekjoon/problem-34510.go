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

    var h1, h2, h3, n int
    fmt.Fscan(reader, &h1, &h2, &h3, &n)
	
	sum := (n/2)*(h2+h3*2)
	if n%2 == 1 {
		sum += (h1+h2+h3)
	}

	fmt.Fprintf(writer, "%v\n", sum)
}
