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

	var n string
	fmt.Fscan(reader, &n)

	if len(n) % 2 == 1 {
		fmt.Fprintf(writer, "%v\n", "fix")
		return
	}
	f := -1
	for i, v := range n {
		if string(v) == "(" {
			f = i
			break
		}
	}
	
	if f == len(n)/2-1 {
		fmt.Fprintf(writer, "%v\n", "correct")
		return
	}
	fmt.Fprintf(writer, "%v\n", "fix")
}
