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

	var n int
	fmt.Fscan(reader, &n)
	
	arr := []int{1, 3, 5, 7, 8, 10, 12}
	for i := 0; i < n; i++ {
		var y, m int
		fmt.Fscan(reader, &y, &m)
		
		isYoon := false
		if y % 4 == 0 && y % 100 != 0 {
			isYoon = true
		}
		if y % 400 == 0 {
			isYoon = true
		}

		ny := y
		nm := m-1
		nd := 30
		if m == 1 {
			ny--
			nm = 12
		}
		
		if isContain(arr, nm) {
			nd = 31
		}
		if m == 3 {
			nd = 28
			if isYoon {
				nd = 29
			}
		}

		fmt.Fprintf(writer, "%v %v %v\n", ny, nm, nd)
	}
	
}

func isContain(arr []int, t int) bool {
	for _, v := range arr {
		if v == t {
			return true
		}
	}
	return false
}
