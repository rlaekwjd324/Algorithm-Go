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

	pre := -11.0
	for {
		var t float64
		fmt.Fscan(reader, &t)

		if t == 999 {
			break
		}
		if pre == -11.0 {
			pre = t
			continue
		}

		v := t-pre
		fmt.Fprintf(writer, "%.2f\n", v)
		pre = t
	}
}
