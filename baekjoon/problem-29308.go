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
	
	r := 0
	rn := ""
	for i := 0; i < n; i++ {
		var v int
		var name, country string
		fmt.Fscan(reader, &v, &name, &country)

		if country == "Russia" {
			if r < v {
				r = v
				rn = name
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", rn)
}
