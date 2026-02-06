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

	for {
		var a, b, c int
		_, err := fmt.Fscan(reader, &a, &b, &c)
		if err != nil {
			return
		}

		cnt := b-a-1
		if b-a < c-b {
			cnt = c-b-1
		}

		fmt.Fprintf(writer, "%v\n", cnt)
	}
}
