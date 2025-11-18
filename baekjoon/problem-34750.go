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
	
	v := 0
	l := 0
	if n >= 1000000 {
		v = int(float64(n)*0.2)
		l = n-v
	} else if n >= 500000 {
		v = int(float64(n)*0.15)
		l = n-v
	} else if n >= 100000 {
		v = int(float64(n)*0.1)
		l = n-v
	} else {
		v = int(float64(n)*0.05)
		l = n-v
	}
	fmt.Fprintf(writer, "%v %v\n", v, l)
	
}
