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
	
	for i := 0; i < t; i++ {
		var n, w, e int
		fmt.Fscan(reader, &n, &w, &e)
		
		sum := 0
		for j := 0; j < n; j++ {
			var lww, lwe, lew, lee int
			fmt.Fscan(reader, &lww, &lwe, &lew, &lee)

			f := lww*w+lew*e
			s := lwe*w+lee*e

			if f > s {
				sum += f
				continue
			}
			sum += s
		}
		fmt.Fprintf(writer, "Data Set %v:\n", i+1)
		fmt.Fprintf(writer, "%v\n\n", sum)
	}
}
