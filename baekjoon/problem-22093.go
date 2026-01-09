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
		var n, a, b int
		fmt.Fscan(reader, &n, &a, &b)
		
		tn := a-b
		if tn < 0 {
			tn = 0
		}
		
		tx := a
		if tx > n-b {
			tx = n-b
		}
		
		fmt.Fprintf(writer, "%v %v\n", tn, tx)
	}
}
