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
		a := make([]int, 5)
		b := make([]int, 5)
		for j := 0; j < 5; j++ {
			var num int
			fmt.Fscan(reader, &num)
			a[j] = num
		}
		for j := 0; j < 5; j++ {
			var num int
			fmt.Fscan(reader, &num)
			b[j] = num

		}

		isYes := false
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				if j != k && a[j] == b[k] {
					isYes = true
					break
				}
			}
			if isYes {
				break
			}
		}
		if isYes {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
