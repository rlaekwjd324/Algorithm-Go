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

	// HEPC
	var a, b string
	fmt.Fscan(reader, &a, &b)

	if a == "HE" && b == "CP" {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}
	if a == "HC" && b == "EP" {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}
	if a == "EP" && b == "HC" {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}
	if a == "EH" && b == "PC" {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}
	if a == "PC" && b == "EH" {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}
	if a == "PE" && b == "CH" {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}
	if a == "CH" && b == "PE" {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}
	if a == "CP" && b == "HE" {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}

	fmt.Fprintf(writer, "%v\n", "NO")
}
