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

	var n int
	fmt.Fscan(reader, &n)
	
	ss := 0.0
	sm := 0.0
	sl := 0.0
	for i := 0; i < n; i++ {
		var s string
		var l float64
		fmt.Fscan(reader, &s, &l)

		switch s {
		case "S":
			ss += l
		case "M":
			sm += l
		case "L":
			sl += l
		}
	}
	
	v := math.Ceil(ss/6.0)+math.Ceil(sm/8.0)+math.Ceil(sl/12.0)
	fmt.Fprintf(writer, "%v\n", v)
}
