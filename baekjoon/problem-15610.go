package main

import (
    "bufio"
    "fmt"
    "os"
	"math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var a float64
	fmt.Fscan(reader, &a)
	
	v := math.Sqrt(a)*4

	fmt.Fprintf(writer, "%v\n", v)
}
