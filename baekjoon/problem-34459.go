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

		v := 5
		c := 1
		for {
			if n <= v {
				a := ""
				if c > 1 {
					a = fmt.Sprintf("^%v", c)
				}
				fmt.Fprintf(writer, "%v%v%v\n", "MIT", a, " time")
				break
			}

			v = v*5
			c++
		}
	}
}
