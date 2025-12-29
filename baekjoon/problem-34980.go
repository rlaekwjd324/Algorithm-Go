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
	var b, a string
	fmt.Fscan(reader, &n, &b, &a)

	if b == a {
		fmt.Fprintf(writer, "%v\n", "Good")
		return
	}
	
	bc := 0
	ac := 0
	for _, v := range b {
		if string(v) == "w" {
			bc++
		}
	}
	for _, v := range a {
		if string(v) == "w" {
			ac++
		}
	}
	if bc == ac {
		fmt.Fprintf(writer, "%v\n", "Its fine")
		return
	}
	if bc < ac {
		fmt.Fprintf(writer, "%v\n", "Manners maketh man")
		return
	}
	fmt.Fprintf(writer, "%v\n", "Oryang")
}
