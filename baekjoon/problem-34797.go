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
	
	sum := 0
	bonus := 0.0
	for i := 0; i < n; i++ {
		var a string
		fmt.Fscan(reader, &a)

		s := 4-int(a[0]-'A')
		t := string(a[1])
		sum += s
		if s >= 2 {
			if t == "1" {
				bonus += 0.05
			} else if t == "2" {
				bonus += 0.025
			}
		}
	}
	
	fmt.Fprintf(writer, "%.6f\n", float64(sum)/float64(n)+bonus)
}
