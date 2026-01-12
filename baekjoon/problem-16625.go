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

	var p, q, s int
	fmt.Fscan(reader, &p, &q, &s)
	
	for i := 1; i <= s; i++ {
		if i % p == 0 && i % q == 0 {
			fmt.Fprintf(writer, "%v\n", "yes")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "no")
}
