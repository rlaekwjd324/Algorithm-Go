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

	for {
		var n int
		fmt.Fscan(reader, &n)
		if n == 0 {
			return
		}
		arr := []string{}
		index := 0
		for i := 0; i < n; i++ {
			var t, m string
			fmt.Fscan(reader, &t, &m)
			switch t {
			case "INSERT":
				if index < len(arr) {
					arr = append(arr[:index], append([]string{m}, arr[index:]...)...)
				} else {
					arr = append(arr, m)
				}
				index++
			case "LEFT":
				index--
				if index < 0 {
					index = 0
				}
			case "RIGHT":
				index++
				if index > len(arr) {
					index = len(arr)
				}
			}
		}
		value := ""
		for i := 0; i < len(arr); i++ {
			value += arr[i]
		}
		fmt.Fprintln(writer, value)
	}
}
