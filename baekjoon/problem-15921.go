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

	var t int
	fmt.Fscan(reader, &t)
	if t == 0 {
		fmt.Fprintf(writer, "%v\n", "divide by zero")
		return
	}
	
	sum := 0.0
	arr := make([]float64, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &arr[i])

		sum += arr[i]
	}

	avg := sum/float64(t)
	
	sum = 0.0
	for i := 0; i < t; i++ {
		sum += (arr[i]/float64(t))
	}

	fmt.Fprintf(writer, "%.2f\n", sum/avg)
}
