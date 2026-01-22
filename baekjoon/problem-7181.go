package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var a string
	var n int
	fmt.Fscan(reader, &a, &n)
	
	for i := 0; i < n; i++ {
		var t string
		fmt.Fscan(reader, &t)
		
		fmt.Fprintf(writer, "%v %v\n", countsOfBall(a, t), countsOfStrike(a, t))
	}
}

func countsOfBall(s, t string) int {
	c := 0
	for _, v := range t {
		if strings.Contains(s, string(v)) {
			c++
		}
	}
	return c
}

func countsOfStrike(s, t string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			c++
		}
	}
	return c
}
