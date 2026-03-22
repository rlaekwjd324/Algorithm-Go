package main

import (
    "bufio"
    "fmt"
    "os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)
	
	min := 0
	for i := 0; i < n-1; i++ {
		if arr[i] == arr[i+1] {
			min = arr[i]*2
			break
		}
	}
	sum := arr[0]+arr[1]

	if min == 0 || (min > 0 && min > sum) {
		min = sum
	}
	fmt.Fprintf(writer, "%v\n", min)
}
