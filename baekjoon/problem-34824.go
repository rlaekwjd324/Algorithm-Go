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
	
	isK := false
	for i := 0; i < n; i++ {
		var t string
		fmt.Fscan(reader, &t)
		
		if t == "korea" {
			isK = true
		}
		if isK && t == "yonsei" {
			fmt.Fprintf(writer, "%v\n", "Yonsei Lost...")
			return
		}
	}
	
	fmt.Fprintf(writer, "%v\n", "Yonsei Won!")
}
