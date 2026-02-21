package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n float64
	fmt.Fscan(reader, &n)

	v := math.Sqrt(n*n - (n/2)*(n/2))

	fmt.Fprintf(writer, "%v\n", v*n/2)
}
