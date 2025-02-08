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

	var n, k int
	fmt.Fscan(reader, &n, &k)

	c := 0
	s := -1
	for i := 0; i < n; i++ {
		var w string
		fmt.Fscan(reader, &w)

		switch w {
		case "save":
			s = c
		case "load":
			if s == -1 {
				c = 0
			} else {
				c = s
			}
		case "shoot":
			if c == 0 {
				continue
			}
			c = c - 1
		case "ammo":
			c += k
		}

		fmt.Fprintln(writer, c)
	}
}
