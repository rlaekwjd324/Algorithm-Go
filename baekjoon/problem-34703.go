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

    var t int
    fmt.Fscan(reader, &t)
	
	arr := make([]bool, 5)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		
		arr[n-1] = true
	}
	for i := 0; i < 5; i++ {
		if !arr[i] {
			fmt.Fprintf(writer, "%v\n", "YES")
			return
		}
 	}
	fmt.Fprintf(writer, "%v\n", "NO")
}
