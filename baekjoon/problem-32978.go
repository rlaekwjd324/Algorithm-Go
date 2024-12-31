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

	arr := make([]string, t)
	for i := 0; i < t; i++ {
		var w string
		fmt.Fscan(reader, &w)

		arr[i] = w
	}
	real := make([]string, t-1)
	for i := 0; i < t-1; i++ {
		var w string
		fmt.Fscan(reader, &w)

		real[i] = w
	}
	for i := 0; i < t; i++ {
		if !isExist(real, arr[i]) {
			fmt.Fprintln(writer, arr[i])
			return
		}
	}
}

func isExist(arr []string, target string) bool {
	for _, a := range arr {
		if a == target {
			return true
		}
	}
	return false
}
