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

	var q int
	fmt.Fscan(reader, &q)
	
	for i := 0; i < q; i++ {
		var a float64
		fmt.Fscan(reader, &a)
		
		isPow := false
		for {
			if a == 1 {
				isPow = true
				break
			}
			if a < 1 {
				break
			}
			a = a/2
		}
		if isPow {
			fmt.Fprintf(writer, "%v\n", 1)
			continue
		}
		fmt.Fprintf(writer, "%v\n", 0)
	}
}
