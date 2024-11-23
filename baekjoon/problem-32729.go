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

	var list string
	fmt.Fscanln(reader, &list)
	listArr := make([]int, 26)
	for i := 0; i < len(list); i++ {
		listArr[list[i]-97]++
	}
	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var word string
		fmt.Fscanln(reader, &word)

		if isContained(listArr, word) {
			fmt.Fprintln(writer, word)
		}
	}
}

func isContained(l []int, w string) bool {
	arr := make([]int, 26)
	for i := 0; i < len(w); i++ {
		arr[w[i]-97]++
	}
	for i := 0; i < 26; i++ {
		if arr[i] > l[i] {
			return false
		}
	}
	return true
}
