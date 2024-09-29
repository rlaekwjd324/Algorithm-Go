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

	for {
		var t int
		fmt.Fscan(reader, &t)
		if t == 0 {
			return
		}

		arr1 := []int{}
		arr2 := []int{}
		for i := 0; i < t; i++ {
			var n int
			fmt.Fscan(reader, &n)
			arr1 = append(arr1, n)
		}
		for i := 0; i < t; i++ {
			var n int
			fmt.Fscan(reader, &n)
			arr2 = append(arr2, n)
		}
		winner := 0
		run1 := 0
		run2 := 0
		count := 0
		for i := 0; i < t; i++ {
			run1 += arr1[i]
			run2 += arr2[i]
			if run1 > run2 {
				if winner == 2 {
					count++
				}
				winner = 1
			}
			if run1 < run2 {
				if winner == 1 {
					count++
				}
				winner = 2
			}
		}
		fmt.Fprintln(writer, count)
	}
}
