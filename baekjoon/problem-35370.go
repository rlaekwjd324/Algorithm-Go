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

	var k, s int
	fmt.Fscan(reader, &k, &s)
	
	q := s / k
	v := s - (k-1) * q	
	
	fmt.Fprintf(writer, "%v\n", v)
}
