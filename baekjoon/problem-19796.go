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

	var m, n int
	fmt.Fscan(reader, &m, &n)
	
	arr := make([]bool, m)
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		
		for j := l-1; j < r; j++ {
			arr[j] = true
		}
	}

	for _, v := range arr {
		if !v {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
	}
	
	fmt.Fprintf(writer, "%v\n", "YES")
}
