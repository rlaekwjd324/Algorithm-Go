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

	var n, m int
	fmt.Fscan(reader, &n, &m)
	
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	for i := 0; i < m; i++ {
		var l, h int
		fmt.Fscan(reader, &l, &h)

		maxValue := indexOfMaximum(arr)
		if arr[h-1] == maxValue {
			continue
		}
		arr[l-1]--
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v ", arr[i])
	}
}

func indexOfMaximum(arr []int) int {
	max := 0
	for _, a := range arr {
		if max < a {
			max = a
		}
	}
	return max
}
