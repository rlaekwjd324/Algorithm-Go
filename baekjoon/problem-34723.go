package main

import (
    "bufio"
    "fmt"
    "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
    defer writer.Flush()

	var P, M, C int64
	fmt.Fscan(reader, &P, &M, &C)

	var X int64
	fmt.Fscan(reader, &X)

	var ans int64 = 1<<63 - 1

	for p := int64(1); p <= P; p++ {
		for m := int64(1); m <= M; m++ {
			pm := p + m
			for c := int64(1); c <= C; c++ {
				value := pm * (m + c)
				diff := abs(value - X)
				if diff < ans {
					ans = diff
				}
			}
		}
	}

	fmt.Fprintln(writer, ans)
}
