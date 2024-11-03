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

	arr := []string{}
	for {
		if n >= 5 {
			arr = append(arr, "3")
			n -= 3
		} else {
			break
		}
	}
	if n == 3 || n == 2 {
		arr = append(arr, fmt.Sprintf("%v", n))
	} else if n == 5 {
		arr = append(arr, "3")
		arr = append(arr, "2")
	} else if n == 4 {
		arr = append(arr, "2")
		arr = append(arr, "2")
	}
	answer := fmt.Sprintf("%v\n%v", len(arr), strings.Join(arr, " "))
	fmt.Fprintln(writer, answer)
}
