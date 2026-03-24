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

	var n, m int
	fmt.Fscan(reader, &n, &m)
	
	v := n+10*m
	
	max := 0
	for i := 0; i < 5; i++ {
		var tn, tm int
		fmt.Fscan(reader, &tn, &tm)
		t := tn+10*tm

		if max < t {
			max = t
		}
	}

	ans := max+1-v
	if ans < 0 {
		ans = 0
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
