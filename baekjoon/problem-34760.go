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

	max := 0
	for i := 0; i < 14; i++ {
		var a int
		fmt.Fscan(reader, &a)
		
		if max < a {
			max = a
		}
	}

	var a int
	fmt.Fscan(reader, &a)
	if max < a {
		fmt.Fprintf(writer, "%v\n", a)
		return
	}
	fmt.Fprintf(writer, "%v\n", max+1)
}
