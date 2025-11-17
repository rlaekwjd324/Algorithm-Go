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

	ds := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for {
		var d, m, y int
		fmt.Fscan(reader, &d, &m, &y)

		if d == 0 && m == 0 && y == 0 {
			break
		}

		sum := d
		for i := 0; i < m-1; i++ {
			sum += ds[i]
		}
		if m > 2 && y % 4 == 0 && (y % 100 != 0 || y % 400 == 0) {
			sum++
		}

		fmt.Fprintf(writer, "%v\n", sum)
	}
}
