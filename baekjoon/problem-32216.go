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

	var n, k, t int
	fmt.Fscan(reader, &n, &k, &t)

	a := 0
	for i := 0; i < n; i++ {
		var d int
		fmt.Fscan(reader, &d)

		var t1 int
		if t > k {
			t1 = t + d - int(math.Abs(float64(t-k)))
		} else if t < k {
			t1 = t + d + int(math.Abs(float64(t-k)))
		} else {
			t1 = t + d
		}
		t = t1
		a += int(math.Abs(float64(t - k)))
	}
	fmt.Fprintln(writer, a)
}
