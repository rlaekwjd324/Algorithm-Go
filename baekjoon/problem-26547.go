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

	var t int
	fmt.Fscan(reader, &t)
	
	for i := 0; i < t; i++ {
		var a string
		fmt.Fscan(reader, &a)
		
		fmt.Fprintf(writer, "%v\n", a)
		if len(a) == 1 {
			continue
		}
		for j := 1; j < len(a)-1; j++ {
			fmt.Fprintf(writer, "%v%v%v\n", string(a[j]), strings.Repeat(" ", len(a)-2), string(a[len(a)-1-j]))
		}
		for j := 0; j < len(a); j++ {
			fmt.Fprintf(writer, "%v", string(a[len(a)-j-1]))
		}
		fmt.Fprintln(writer)
	}
	
}
