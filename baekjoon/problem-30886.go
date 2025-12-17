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

	var a float64
	fmt.Fscan(reader, &a)
	
	// a = pie*r*r
	r := math.Sqrt(a/math.Pi)
	v := (r*2+2)*(r*2+2)

	fmt.Fprintf(writer, "%.9f\n", v)
}
