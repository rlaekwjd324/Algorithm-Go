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

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscanln(reader, &num)

		l := ""
		if num%2 == 1 {
			l += "O"
		}
		if isSquareNumber(num) {
			l += "S"
		}
		if l == "" {
			l = "EMPTY"
		}
		fmt.Fprintf(writer, "%s\n", l)
	}
}

func isSquareNumber(num int) bool {
	for i := 0; i <= 1000; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
