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
		var n, m int
		fmt.Fscan(reader, &n, &m)

		nArr := make([]int, n)
		mArr := make([]int, m)
		
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &nArr[j])
		}
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &mArr[j])
		}

		if n <= m {
			fmt.Fprintf(writer, "%v\n", "Yes")
			continue
		}
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
