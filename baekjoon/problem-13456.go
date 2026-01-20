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
		
		uArr := make([]int, n)
		c := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &uArr[j])
		}
		for j := 0; j < n; j++ {
			var a int
			fmt.Fscan(reader, &a)

			if a != uArr[j] {
				c++
			}
		}

		fmt.Fprintf(writer, "%v\n", c)
	}	
}
