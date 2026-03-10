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
	var s string
	fmt.Fscan(reader, &n, &s)
	
	counts := make([]int, 5)
	hiarc := []string{"H", "I", "A", "R", "C"}

	for _, v := range s {
		for i, j := range hiarc {
			if j == string(v) {
				counts[i]++
			}
		}
	}

	min := n
	for _, v := range counts {
		if min > v {
			min = v
		}
	}

	fmt.Fprintf(writer, "%v\n", min)
}
