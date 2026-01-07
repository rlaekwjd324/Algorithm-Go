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
		var n int
		fmt.Fscan(reader, &n)

		if n == 0 {
			break
		}

		v := (n*(n+1)/2)*(n*(n+1)/2)
		
		fmt.Fprintf(writer, "%v\n", v)
	}
}
