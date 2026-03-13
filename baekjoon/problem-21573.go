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

	var b, a int
	var m string
	fmt.Fscan(reader, &b, &a, &m)
	
	switch m {
	case "freeze":
		if b > a {
			b = a
		}
		fmt.Fprintf(writer, "%v\n", b)
		return
	case "heat":
		if b < a {
			b = a
		}
		fmt.Fprintf(writer, "%v\n", b)
		return
	case "auto":
		fmt.Fprintf(writer, "%v\n", a)
		return
	case "fan":
		fmt.Fprintf(writer, "%v\n", b)
		return
	}
}
