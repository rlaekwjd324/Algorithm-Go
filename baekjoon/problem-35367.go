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

	var r, c int
	fmt.Fscan(reader, &r, &c)
	
	arr := make([]string, c)
	for i := 0; i < r; i++ {
		var a string
		fmt.Fscan(reader, &a)
		
		for j, v := range a {
			if string(v) == "." {
				continue
			}
			arr[j] = string(v)
		}
	}
	
	fmt.Fprintf(writer, "%v\n", strings.Join(arr, ""))
}
