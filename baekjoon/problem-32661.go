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
	fmt.Fscanln(reader, &n)

	min := 100000
	max := 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscanln(reader, &num)

		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	if max/2 > min {
		fmt.Fprintln(writer, 0)
		return
	}
	fmt.Fprintln(writer, min-max/2)
}
