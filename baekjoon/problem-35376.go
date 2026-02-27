package main

import (
    "bufio"
    "fmt"
    "os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

	var c int
	fmt.Fscan(reader, &c)
	
	arr := make([]int, c)
	sum := 0
	for i := 0; i < c; i++ {
		fmt.Fscan(reader, &arr[i])
		sum += arr[i]
	}
	sort.Ints(arr)

	cnt := arr[c-2]
	r := 0
	for i, v := range arr {
		if i == c-2 {
			fmt.Fprintf(writer, "%v\n", "IMPOSSIBLE TO WIN")
			return
		}
		cnt += v
		r++
		if sum/2 < cnt {
			fmt.Fprintf(writer, "%v\n", r)
			return
		}
	}
	
}
