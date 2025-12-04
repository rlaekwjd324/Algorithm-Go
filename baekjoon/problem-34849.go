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
	
	if n > 10000 {
		fmt.Fprintf(writer, "%v\n", "Time limit exceeded")
		return
	}
	fmt.Fprintf(writer, "%v\n", "Accepted")
}
