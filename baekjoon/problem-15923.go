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

	var n int
	fmt.Fscan(reader, &n)
	
	preX := -1.0
	preY := -1.0
	firstX := -1.0
	firstY := -1.0
	sum := 0.0
	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Fscan(reader, &x, &y)

		if i == 0 {
			firstX = x
			firstY = y
		}
		if i > 0 {
			sum += math.Abs(preX-x)
			sum += math.Abs(preY-y)
		}
		preX = x
		preY = y
	}
	sum += math.Abs(preX-firstX)
	sum += math.Abs(preY-firstY)
	
	fmt.Fprintf(writer, "%v\n", sum)
}
