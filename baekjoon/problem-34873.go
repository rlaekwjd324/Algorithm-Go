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
	
	arr := make([]int, n*2)
	for i := 0; i < n*2; i++ {
		var a int
		fmt.Fscan(reader, &a)
		arr[a-1]++
		if arr[a-1] > 2 {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}
	
	fmt.Fprintf(writer, "%v\n", "Yes")
}
