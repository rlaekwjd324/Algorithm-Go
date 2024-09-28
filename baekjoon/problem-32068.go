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
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var l, r, s int
		fmt.Fscan(reader, &l, &r, &s)

		ls := s - l
		rs := r - s
		if rs <= ls {
			fmt.Fprintln(writer, rs*2)
		}else {
			fmt.Fprintln(writer, ls*2+1)
		}
	}
}
