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
		var a, b string
		fmt.Fscan(reader, &a, &b)
		
		if a == b {
			fmt.Fprintf(writer, "%v\n", "OK")
			continue
		}
		fmt.Fprintf(writer, "%v\n", "ERROR")
	}
}
