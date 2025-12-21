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

	c := 1
	for {
		var n int
		fmt.Fscan(reader, &n)
		
		if n == 0 {
			return
		}
		
		v := 0.0
		o := 0.0
		for	i := 0; i < n; i++ {
			var c, b, l, ni float64
			fmt.Fscan(reader, &c, &b, &l, &ni)

			kg := c/85.0+b/85.0+l/85.0
			v += c*0.8+b*1.0+l*1.2+ni*0.8
			o += c/85.0*7.5+b/85.0*24.0+l/85.0*32.0+ni*0.2+kg*8.0
		}

		fmt.Fprintf(writer, "Case #%v: RM%.2f\n", c, v-o)
		c++
	}
}
