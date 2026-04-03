package main

import (
    "bufio"
    "fmt"
    "os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)
	
	arr := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		t := s[i:]
		arr[i] = t
	}

	sort.Strings(arr)
	
	fmt.Fprintf(writer, "%v\n", arr[len(s)-1])
}
