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

	var a, b, c, d int
	fmt.Fscan(reader, &a, &b, &c, &d)
	
	sh, w := false, false

	if a+b <= d {
		sh = true
	}
	if c <= d {
		w = true
	}

	if sh && w {
		fmt.Fprintf(writer, "%v\n", "~.~")
		return
	}
	if sh {
		fmt.Fprintf(writer, "%v\n", "Shuttle")
		return
	}
	if w {
		fmt.Fprintf(writer, "%v\n", "Walk")
		return
	}
	fmt.Fprintf(writer, "%v\n", "T.T")
}
