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
	
	max := 1
	c := 1
	pre := 0
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		
		if i == 0 {
			pre = n
			continue
		}
		if pre < n {
			pre = n
			c++
			if max < c {
				max = c
			}
		} else {
			pre = n
			c = 1
		}
	}

	fmt.Fprintf(writer, "%v\n", max)
}
