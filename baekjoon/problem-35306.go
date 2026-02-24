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

	var n, k int
	fmt.Fscan(reader, &n, &k)
	
	maxArr := make([]int, k)
	idxArr := make([]int, k)
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			var a int
			fmt.Fscan(reader, &a)
			
			if maxArr[j] < a {
				maxArr[j] = a
				idxArr[j] = i
			} else if maxArr[j] == a {
				idxArr[j] = -1
			}
		}
	}
	
	cnt := 0
	for i := 0; i < n; i++ {
		for _, v := range idxArr {
			if v == i {
				cnt++
				break
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
