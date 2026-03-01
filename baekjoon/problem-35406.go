package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)
	
	ha := strings.Count(s, "ha")
	boooo := strings.Count(s, "boooo")
	bravo := strings.Count(s, "bravo")

	fmt.Fprintf(writer, "%v\n", ha - boooo + bravo * 3)
}
