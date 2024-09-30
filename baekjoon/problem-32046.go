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

	for {
		var t int
		fmt.Fscan(reader, &t)
		
		if t == 0 {
			return
		}

		sum := 0
		for i := 0; i < t; i++ {
			var n int
			fmt.Fscan(reader, &n)
			if sum+n > 300 {
				continue
			}
			sum = sum + n
		}
		fmt.Fprintln(writer, sum)
	}
}
