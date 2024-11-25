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

	var w, s int
	fmt.Fscanln(reader, &w, &s)

	c := (s + 1) * s / 2

	for i := 0; i <= c; i++ {
		if i*29260+(c-i)*29370 == w {
			fmt.Fprintln(writer, c-i)
			return
		}
	}
}
