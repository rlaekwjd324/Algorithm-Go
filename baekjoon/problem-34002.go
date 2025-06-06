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

	var a, b, c, s, v float64
	var l int
	fmt.Fscan(reader, &a, &b, &c, &s, &v, &l)

	sum := float64((250 - l) * 100)

	if v*c*30 >= sum {
		fmt.Fprintf(writer, "%v\n", math.Ceil(sum/c))
		return
	}

	cnt := 30.0 * v
	sum -= v * c * 30.0
	if s*b*30 >= sum {
		fmt.Fprintf(writer, "%v\n", math.Ceil(sum/b+cnt))
		return
	}

	sum -= s * b * 30.0
	cnt += (30.0 * s)
	fmt.Fprintf(writer, "%v\n", math.Ceil(sum/a+cnt))
}
