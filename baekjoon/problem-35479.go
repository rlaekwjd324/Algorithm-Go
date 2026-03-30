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

	var r, g, b float64
	fmt.Fscan(reader, &r, &g, &b)
	
	nr := r/255.0
	ng := g/255.0
	nb := b/255.0

	max := math.Max(nr, ng)
	max = math.Max(max, nb)
	k := 1-max

	if k == 1 {
		fmt.Fprintf(writer, "%v %v %v %v\n", 0, 0, 0, k)
		return
	}

	c := (1-nr-k)/(1-k)
	m := (1-ng-k)/(1-k)
	y := (1-nb-k)/(1-k)

	fmt.Fprintf(writer, "%v %v %v %v\n", c, m, y, k)
}
