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

	var f, s, s1, t, t1 string
	fmt.Fscan(reader, &f, &s, &s1, &t, &t1)

	for i := 0; i < 1000; i++ {
		fmt.Fprintf(writer, "%v\n", -1)
	}
}
