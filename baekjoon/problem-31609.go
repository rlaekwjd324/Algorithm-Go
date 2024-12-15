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

	arr := make([]bool, 10)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(reader, &num)
		arr[num] = true
	}
	for i := 0; i < 10; i++ {
		if arr[i] {
			fmt.Fprintln(writer, i)
		}
	}
}
