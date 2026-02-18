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
	
	cArr := []int{8, 4, 2, 1}
	lArr := []int{209, 219, 229, 239}

	max := 0
	idx := -1
	for i := 0; i < len(cArr); i++ {
		t := n + cArr[i]
		if n <= lArr[i] && max <= t {
			if t > lArr[i]+1 {
				t = lArr[i]+1
			}
			max = t
			idx = i+1
		}
	}
	fmt.Fprintf(writer, "%v\n", idx)
}
