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

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if k == 0 || k == n-1 {
					fmt.Fprintf(writer, "%v", "#")
					continue
				}
				if j > n/2 || j == 0 {
					fmt.Fprintf(writer, "%v", ".")
					continue
				}
				if j == n/2 {
					if k == n/2 {
						fmt.Fprintf(writer, "%v", "#")
						continue
					}
					fmt.Fprintf(writer, "%v", ".")
					continue
				}
				if k == j || k == n-j-1 {
					fmt.Fprintf(writer, "%v", "#")
					continue
				}
				fmt.Fprintf(writer, "%v", ".")
				continue
			}
			fmt.Fprintln(writer)
		}
	}
}
