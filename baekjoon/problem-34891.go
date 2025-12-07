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

	var n, m float64
	fmt.Fscan(reader, &n, &m)
	
	fmt.Fprintf(writer, "%v\n", math.Ceil(n/m))
}
