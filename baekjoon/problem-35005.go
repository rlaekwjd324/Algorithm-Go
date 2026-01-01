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

	var h, m int
	fmt.Fscan(reader, &h, &m)

	c := 0
	for i := 0; i < h; i++ {
		strI := fmt.Sprintf("%v", i)
		for j := 0; j < m; j++ {
			strJ := fmt.Sprintf("%v", j)
			if j == 39 {
				if string(strI[len(strI)-1]) == "2" {
					c++
					continue
				}
				if i == 20 || i == 22 || i == 23 {
					c++
					continue
				}
			}
			if i == 23 {
				if string(strJ[0]) == "9" {
					c++
					continue
				}
				if j == 9 || j == 39 || j == 99 {
					c++
					continue
				}
			}
		}
	}
	
	fmt.Fprintf(writer, "%v\n", c)
}
