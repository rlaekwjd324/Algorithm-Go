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

	var n, k int
	fmt.Fscan(reader, &n)
	
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	fmt.Fscan(reader, &k)

	for i := 0; i < k; i++ {
		var p int
		fmt.Fscan(reader, &p)
		
		arr[p-1]--
	}

	for i := 0; i < n; i++ {
		if arr[i] < 0 {
			fmt.Fprintf(writer, "%v\n", "yes")
			continue
		}
		fmt.Fprintf(writer, "%v\n", "no")
	}
}
