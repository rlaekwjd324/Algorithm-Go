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
   
	var n int
	fmt.Fscan(reader, &n)

	c := -1
	
	for {
		c++
		if n == 1 {
			break
		}
		if n % 2 == 0 {
			n = n/2
			continue
		}
		n += n*2+1
	}
	
	fmt.Fprintf(writer, "%v\n", c)
}
