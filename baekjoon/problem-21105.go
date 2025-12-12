package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var p, c float64
		fmt.Fscan(reader, &p, &c)
		// C=100%(P−O​)/O
		// OC = 100%(P-O)
		// OC/100 + O = P
		// (C/100+1)O = P
		o := p / (c/100 + 1)
		fmt.Fprintf(writer, "%.6f\n", o)
	}

}
