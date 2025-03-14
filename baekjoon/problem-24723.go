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

	var n int
	fmt.Fscan(reader, &n)

	fmt.Fprintln(writer, int(math.Pow(2.0, float64(n))))
}
