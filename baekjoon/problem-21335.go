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
	
	// a = pi*r*r
	r := math.Sqrt(a/math.Pi)
	// answer := 2*pi*r
	
	fmt.Fprintf(writer, "%v\n", 2*math.Pi*r)
}
