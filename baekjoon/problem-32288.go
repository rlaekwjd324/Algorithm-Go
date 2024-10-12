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

	var n int
	fmt.Fscanln(reader, &n)

	var nickname string
	fmt.Fscanln(reader, &nickname)
	arr := strings.Split(nickname, "")

	for i := 0; i < n; i++ {
		if arr[i] == "l" {
			arr[i] = "L"
			continue
		}
		if arr[i] == "I" {
			arr[i] = "i"
			continue
		}
	}

	fmt.Fprintln(writer, strings.Join(arr, ""))
}
