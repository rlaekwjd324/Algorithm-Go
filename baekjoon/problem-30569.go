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

	var h1, d1, t1, h2, d2, t2 int
	fmt.Fscan(reader, &h1, &d1, &t1, &h2, &d2, &t2)

	c := 0
	l1, l2 := false, false
	for {
		if h1 <= 0 {
			l1 = true
		}
		if h2 <= 0 {
			l2 = true
		}
		if l1 || l2 {
			break
		}
		if c % t1 == 0 {
			h2 -= d1
		}
		if c % t2 == 0 {
			h1 -= d2
		}
		c++
	}
	
	if l1 && l2 {
		fmt.Fprintf(writer, "%v\n", "draw")
		return
	}
	if l1 {
		fmt.Fprintf(writer, "%v\n", "player two")
		return
	}
	fmt.Fprintf(writer, "%v\n", "player one")
}
