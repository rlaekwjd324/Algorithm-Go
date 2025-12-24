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
	fmt.Fscan(reader, &n)

	if n+2 >= 8 {
		fmt.Fprintf(writer, "%v\n", "Success!")
		return
	}
		
	fmt.Fprintf(writer, "%v\n", "Oh My God!")
}
