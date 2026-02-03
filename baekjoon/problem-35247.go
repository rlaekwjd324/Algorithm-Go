package main

import (
    "bufio"
    "fmt"
    "os"
	"math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var t int64
	fmt.Fscan(reader, &t)
	
	c := 1
	for {
		if math.Pow(2.0, float64(c))-1 >= float64(t) {
			if c == 1 {
				fmt.Fprintf(writer, "%v bit\n", c)
				return
			}
			fmt.Fprintf(writer, "%v bits\n", c)
			return
		}
		c = c*2
	}

}
