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

	var a big.Int
	var b string
	fmt.Fscan(reader, &a, &b)
	fmt.Fprintf(writer, "%v\n", a.String())
	fmt.Fprintf(writer, "%v\n", b)

	sum := new(big.Int)
	c := big.NewInt(1)
	for i, _ := range b {
		m := int64(b[len(b)-i-1]-'0')
		v := new(big.Int)
		v.Mul(big.NewInt(m), &a)
		fmt.Fprintf(writer, "%v\n", v)

		s := new(big.Int)
		s.Mul(v, c)
		sum.Add(sum, s)
		c.Mul(c, big.NewInt(10))
	}

	fmt.Fprintf(writer, "%v\n", sum)
}
