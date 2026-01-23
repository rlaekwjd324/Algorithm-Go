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

	var n, r, s int
	fmt.Fscan(reader, &n, &r, &s)

	for i := 0; i < n; i++ {
		arr := make([]int, s)
		for j := 0; j < r; j++ {
			var t string
			fmt.Fscan(reader, &t)

			for k, v := range t {
				if string(v) == "#" {
					arr[k] = j + 1
				}
			}
		}
		max := 0
		min := r
		for j := 0; j < s; j++ {
			if arr[j] > max {
				max = arr[j]
			}
			if arr[j] < min {
				min = arr[j]
			}
		}
		fmt.Fprintf(writer, "%v\n", max-min)
	}

}
