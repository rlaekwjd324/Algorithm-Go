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

	var k, n int
	fmt.Fscan(reader, &k, &n)

	arr := make([]string, n)
	for i := 0; i < n-1; i++ {
		arr[i] = fmt.Sprintf("%v", i+1)
		k -= (i + 1)
	}
	arr[n-1] = fmt.Sprintf("%v", k)

	fmt.Fprintln(writer, strings.Join(arr, "\n"))
}
