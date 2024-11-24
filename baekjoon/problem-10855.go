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

	pre := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		if pre > a {
			fmt.Fprintln(writer, "no")
			return
		}
		pre = a
	}
	fmt.Fprintln(writer, "yes")
}
