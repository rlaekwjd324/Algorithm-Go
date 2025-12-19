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

	var arr = make([]int, 3)
	var c1, c2, c3 float64
	fmt.Fscan(reader, &arr[0], &arr[1], &arr[2], &c1, &c2, &c3)

	t := float64(arr[0]+arr[1]+arr[2]) * (c1) / 100.0
	sort.Ints(arr)
	var maxC, minC float64
	if c2 < c3 {
		maxC = c3
		minC = c2
	} else {
		maxC = c2
		minC = c3
	}
	s := float64(arr[1])*minC/100.0 + float64(arr[2])*maxC/100.0

	if t > s {
		fmt.Fprintf(writer, "%v %.2f\n", "one", t)
		return
	}
	fmt.Fprintf(writer, "%v %.2f\n", "two", s)
}
