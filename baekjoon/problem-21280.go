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
	
	arr := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	
	p := arr[0]
	c := 0
	q := 0
	for i := 1; i < t; i++ {
		if p < arr[i] {
			c += arr[i]-p	
		}
		if p > arr[i] {
			q += p-arr[i]
		}
		p = arr[i]
	}

	fmt.Fprintf(writer, "%v %v\n", q, c)
}
