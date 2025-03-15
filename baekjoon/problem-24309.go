package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var a, b, c big.Int
	fmt.Fscan(reader, &a, &b, &c)

	d := new(big.Int)
	d = d.Sub(&b, &c)
	d = d.Div(d, &a)
	fmt.Fprintln(writer, d)
}
