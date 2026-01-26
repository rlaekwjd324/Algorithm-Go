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

	var a, b int
	fmt.Fscan(reader, &a, &b)
	
	if a%2 == 0 || b%2 == 0 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	if a > b {
		fmt.Fprintf(writer, "%v\n", b)
		return
	}

	fmt.Fprintf(writer, "%v\n", a)
}
