package main

import (
    "bufio"
    "fmt"
    "os"
	"math/big"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var n big.Int
	fmt.Fscan(reader, &n)
	
	mul := new(big.Int)
	mul.Mul(&n, big.NewInt(2))
	
	fmt.Fprintf(writer, "%v\n", mul)
}
