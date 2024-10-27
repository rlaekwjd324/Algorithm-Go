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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	sum := 0
	index := -1
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscanln(reader, &a)
		sum += a
		if sum >= m {
			index = i + 1
			break
		}
	}
	fmt.Fprintln(writer, index)
}
