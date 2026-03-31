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
		var n, s int
		fmt.Fscan(reader, &n, &s)

		v := 10000000+n

		if v == s {
			fmt.Fprintf(writer, "%v\n", "Yes")
			continue
		}
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
