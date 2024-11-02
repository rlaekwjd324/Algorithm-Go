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

	sum := 1
	arr := make([]int, n)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = arr[i-1] + i + 1
		sum += arr[i]
	}
	fmt.Fprintln(writer, sum)
}
