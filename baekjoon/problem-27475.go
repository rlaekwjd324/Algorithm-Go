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
	
	for i := 0; i < n; i++ {
		var b, l int
		fmt.Fscan(reader, &b, &l)
		
		c := 0
		bArr := make([]bool, 100)
		for j := 0; j < b; j++ {
			var t int
			fmt.Fscan(reader, &t)
			
			bArr[t-1] = true
		}
		for j := 0; j < l; j++ {
			var t int
			fmt.Fscan(reader, &t)
			
			if bArr[t-1] {
				c++
			}
		}
		
		fmt.Fprintf(writer, "%v\n", c)
	}
}
