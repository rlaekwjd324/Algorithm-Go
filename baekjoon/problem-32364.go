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
	fmt.Fscan(reader, &n)

	var arr []int
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		arr = append(arr, a)
	}
	count := 0
	for i := 0; i < n/2; i++ {
		if arr[i] == arr[n/2+i] {
			count += 2
		}
	}
	fmt.Fprintln(writer, count)
}
