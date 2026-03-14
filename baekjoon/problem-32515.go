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
	var j1, j2, t1, t2 string
	fmt.Fscan(reader, &n, &j1, &j2, &t1, &t2)
	
	new := ""
	for i := range j1 {
		if j1[i] == t1[i] {
			if j2[i] != t2[i] {
				fmt.Fprintf(writer, "%v\n", "htg!")
				return
			}
			new += string(j2[i])
		}
	}

	fmt.Fprintf(writer, "%v\n", new)
}
