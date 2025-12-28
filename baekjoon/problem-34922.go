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

	var w, h, r float64
	fmt.Fscan(reader, &w, &h, &r)
	
	v := w*h-math.Pi*r*r/4.0
	
	fmt.Fprintf(writer, "%.2f\n", v)
}
