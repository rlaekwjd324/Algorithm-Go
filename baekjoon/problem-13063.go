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

	for {
		var n, m, k float64
		fmt.Fscan(reader, &n, &m, &k)
		if n == 0 && m == 0 && k == 0 {
			return
		}

		if m > math.Ceil(n/2.0) {
			fmt.Fprintf(writer, "%v\n", 0)
			continue
		}
		d := n - m - k
		if math.Floor(n/2.0)+1-m > d {
			fmt.Fprintf(writer, "%v\n", -1)
			continue
		}
		fmt.Fprintf(writer, "%v\n", math.Floor(n/2.0)+1-m)
	}
	
	
}
