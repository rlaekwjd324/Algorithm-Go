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
	
	switch string(n[0]) {
	case "F":
		fmt.Fprintf(writer, "%v\n", "Foundation")
	case "C":
		fmt.Fprintf(writer, "%v\n", "Claves")
	case "V":
		fmt.Fprintf(writer, "%v\n", "Veritas")
	case "E":
		fmt.Fprintf(writer, "%v\n", "Exploration")
	}
	
}
