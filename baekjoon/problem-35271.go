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

	arr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	for i := 0; i < n; i++ {
		tarr := make([]int, 3)
		fmt.Fscan(reader, &tarr[0])
		fmt.Fscan(reader, &tarr[1])
		fmt.Fscan(reader, &tarr[2])

		for i := 0; i < 3; i++ {
			if arr[tarr[i]-1] > 0 {
				fmt.Fprintf(writer, "%v ", tarr[i])
				arr[tarr[i]-1]--
				break
			}
			if i == 2 {
				fmt.Fprintf(writer, "%v ", -1)
			}
		}
	}
}
