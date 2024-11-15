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

	answer := "AKARAKA"
	for i := 1; i < n; i++ {
		answer += "RAKA"
	}
	fmt.Fprintln(writer, answer)
}
