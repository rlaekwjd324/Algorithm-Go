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

	var a1, b1, c1, a2, b2, c2, h1, m1, s1 int
	fmt.Fscan(reader, &a1, &b1, &c1, &a2, &b2, &c2, &h1, &m1, &s1)
	
	v1 := h1*b1*c1+m1*c1+s1
	vh := v1/c2/b2
	v1 = v1-vh*c2*b2
	vm := v1/c2
	
	fmt.Fprintf(writer, "%v %v %v\n", vh, vm, v1%c2)
}
