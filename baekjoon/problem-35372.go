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
	
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	
	for i := 1; i <= 28; i++ {
		for j := 0; j < n; j++ {
			if getMoon(i*(j+1)) != arr[j] {
				break
			}
			if j == n-1 {
				fmt.Fprintf(writer, "%v\n", i)
				return
			}
		}
	}
}

func getMoon(d int) string {
	m := d % 28
	if m == 14 {
		return "Full"
	}
	if m == 0 {
		return "New"
	}
	if m >= 23 || m <= 4 {
		return "Crescent"
	}
	if m >= 20 || m <= 8 {
		return "Quarter"
	}
	return "Gibbous"
}
