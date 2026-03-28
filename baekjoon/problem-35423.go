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

	var s string
	fmt.Fscan(reader, &s)

	arr := make([]int, 2)
	for _, v := range s {
		switch string(v) {
		case "S":
			arr[0]++
		case "L":
			arr[1]++
		}
	}

	fmt.Fprintf(writer, "%v%v\n", strings.Repeat("SciCom", arr[0]), strings.Repeat("Love", arr[1]))
}
