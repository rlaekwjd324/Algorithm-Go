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

	var m, d float64
	fmt.Fscan(reader, &m, &d)
	
	fmt.Fprintf(writer, "%v\n", math.Ceil(d/m))
}
