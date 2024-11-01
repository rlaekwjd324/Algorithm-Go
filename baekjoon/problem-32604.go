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

	a := -1
	b := -1
	for i := 0; i < n; i++ {
		var tempA, tempB int
		fmt.Fscanln(reader, &tempA, &tempB)
		if a > tempA || b > tempB {
			fmt.Fprintln(writer, "no")
			return
		}
		a = tempA
		b = tempB
	}
	fmt.Fprintln(writer, "yes")
}
