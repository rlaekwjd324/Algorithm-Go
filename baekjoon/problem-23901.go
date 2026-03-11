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
		
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &arr[j])
		}

		count := 0
		for j := 1; j < n-1; j++ {
			if arr[j] > arr[j-1] && arr[j] > arr[j+1] {
				count++
			}
		}

		fmt.Fprintf(writer, "Case #%v: %v\n", i+1, count)
	}
}
