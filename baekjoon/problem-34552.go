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

	arr := make([]int, 11)
	for i := 0; i < 11; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	var n int
	fmt.Fscan(reader, &n)

	sum := 0
	for i := 0; i < n; i++ {
		var b, s int
		var l float64
		fmt.Fscan(reader, &b, &l, &s)

		if l >= 2.0 && s >= 17 {
			sum += arr[b]
		}
	}

	fmt.Fprintf(writer, "%v\n", sum)
}
