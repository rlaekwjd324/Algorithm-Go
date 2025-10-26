package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
   
	var a, b, x, y int
	fmt.Fscan(reader, &a, &b, &x, &y)

	lax := a
	if lax < x {
		lax = x
	}
	v1 := b+y+lax
	lby := b
	if lby < y {
		lby = y
	}
	v2 := a+x+lby

	if v1 > v2 {
		fmt.Fprintf(writer, "%v\n", v2*2)
		return
	}
	fmt.Fprintf(writer, "%v\n", v1*2)
}
