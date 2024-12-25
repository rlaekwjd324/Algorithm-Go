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

	var n, m int
	fmt.Fscan(reader, &n, &m)

	var sArr []string
	var cArr []int
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(reader, &k)
		for j := 0; j < k; j++ {
			var s string
			fmt.Fscan(reader, &s)

			index := IndexOf(sArr, s)
			if index == -1 {
				sArr = append(sArr, s)
				cArr = append(cArr, 1)
				continue
			}
			cArr[index]++
		}
	}
	count := 0
	for _, c := range cArr {
		if c >= m {
			count++
		}
	}

	fmt.Fprintln(writer, count)
}

func IndexOf(array []string, target string) int {
	for index, value := range array {
		if value == target {
			return index
		}
	}
	return -1
}
