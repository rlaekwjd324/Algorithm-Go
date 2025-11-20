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
	
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			var a int
			fmt.Fscan(reader, &a)
			arr[i][j] = a
		}
	}
	
	sum := 0
	for i := 0; i < m; i++ {
		max := 0
		for j := 0; j < n; j++ {
			if max < arr[j][i] {
				max = arr[j][i]
			}
		}
		sum += max
	}
	
	fmt.Fprintf(writer, "%v\n", sum)
}
