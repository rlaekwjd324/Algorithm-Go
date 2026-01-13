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
	
	arr := make([]int, 3)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			var r int
			fmt.Fscan(reader, &r)
			
			arr[r-1]++
		}
	}

	f := 0
	l := 0
	for i := 0; i < 3; i++ {
		if arr[i] < n {
			f = i+1
		} 
		if arr[i] > n {
			l = i+1
		} 
	}
	fmt.Fprintf(writer, "%v\n%v\n", f, l)
}
