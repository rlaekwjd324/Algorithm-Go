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

	var a, b int
	fmt.Fscan(reader, &a, &b)
	var d int
	fmt.Fscan(reader, &d)

	count := math.Ceil(float64(a*b) / 12.0)
	value := int(count) * d

	fmt.Fprintln(writer, value)
}
