package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
   
	var n, v, t int
	fmt.Fscan(reader, &n, &v, &t)
	
	for i := 0; i < t; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		
		if v == a {
			v = b
			continue
		}
		if v == b {
			v = a
		}
	}

	fmt.Fprintf(writer, "%v\n", v)
}
