package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var n int
	var s string
	fmt.Fscan(reader, &n, &s)

	arr := make([]int, 10)

	for _, v := range s {
		switch string(v) {
		case "L":
			for j := 0; j < 10; j++ {
				if arr[j] == 0 {
					arr[j] = 1
					break
				}
			}
		case "R":
			for j := 9; j >= 0; j-- {
				if arr[j] == 0 {
					arr[j] = 1
					break
				}
			}
		default:
			i, _ := strconv.Atoi(string(v))
			arr[i] = 0
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Fprintf(writer, "%v", arr[i])
	}
}
