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

	var t int
	fmt.Fscan(reader, &t)
	
	for i := 0; i < t; i++ {
		var a1, p1, r1, p2 float64
		fmt.Fscan(reader, &a1, &p1, &r1, &p2)

		v1 := a1/p1
		v2 := (r1*r1*math.Pi)/p2
		
		if v1 > v2 {
			fmt.Fprintf(writer, "%v\n", "Slice of pizza")
			continue
		}
		fmt.Fprintf(writer, "%v\n", "Whole pizza")
	}
}
