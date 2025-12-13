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
		var s string
		fmt.Fscan(reader, &s)

		ans := 0
		c := 0
		for i, v := range s {
			if i%8 == 0 {
				c = 0
			}
			if i%8 == 7 {
				if (c%2 == 0 && string(v) == "0") || (c%2 == 1 && string(v) == "1") {
					continue
				}
				ans++
				continue
			}
			if string(v) == "1" {
				c++
			}
		}
		fmt.Fprintf(writer, "%v\n", ans)
	}

}
