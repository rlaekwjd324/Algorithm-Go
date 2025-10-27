package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
   
	var n, m int
	fmt.Fscan(reader, &n, &m)
	
	arr := make([][]string, n)
	for i := 0; i < n; i++ {
		var p string
		fmt.Fscan(reader, &p)

		arr[i] = strings.Split(p, "")
	}

	c := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch arr[i][j] {
			case "X":
				if i > 0 && arr[i-1][j] == "Y" {
					c++
				}
				if i < n-1 && arr[i+1][j] == "Y" {
					c++
				}
				if j > 0 && arr[i][j-1] == "Y" {
					c++
				}
				if j < m-1 && arr[i][j+1] == "Y" {
					c++
				}
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", c)
}
