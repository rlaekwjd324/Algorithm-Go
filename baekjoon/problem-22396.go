package main

import (
    "bufio"
    "fmt"
    "os"
	"math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	for {
		var r0, w0, c, r float64
		fmt.Fscan(reader, &r0, &w0, &c, &r)

		if r0 == 0.0 && w0 == 0.0 && c == 0.0 && r == 0.0 {
			return
		}
		
		// (r0+x*r)/w0 == c
		// r0+x*r = c*w0
		// x*r = c*w0-r0
		x := (c*w0-r0)/r
		if x < 0 {
			x = 0.0
		}

		fmt.Fprintf(writer, "%v\n", math.Ceil(x))
	}
}
