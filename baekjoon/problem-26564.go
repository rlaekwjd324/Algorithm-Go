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
	fmt.Fscan(reader, &n)
	
	arr := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
	for i := 0; i < n; i++ {
		counts := make([]int, 13)
		max := 0
		for j := 0; j < 5; j++ {
			var a string
			fmt.Fscan(reader, &a)
			
			idx := indexOf(arr, string(a[0]))
			counts[idx]++
			if max < counts[idx] {
				max = counts[idx]
			}
		}
		
		fmt.Fprintf(writer, "%v\n", max)
	}	
}

func indexOf(arr []string, t string) int {
	for i, v := range arr {
		if v == t {
			return i
		}
	}
	return -1
}
