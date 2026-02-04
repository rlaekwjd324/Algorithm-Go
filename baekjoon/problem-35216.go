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
	
	a := t/2
	b := t/2+t%2

	for {
		if a * b == 0 {
			a--
			b++
		} else {
			break
		}
	}
	
	fmt.Fprintf(writer, "%v %v\n", a, b)
}
