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

	var n int
	var s string
	fmt.Fscan(reader, &n, &s)

	arr := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII"}
	for i, v := range arr {
		if strings.Contains(s, v) {
			fmt.Fprintf(writer, "%v ", i+1)
		}
	}
}
