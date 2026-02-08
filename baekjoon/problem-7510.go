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

	var t int
	fmt.Fscan(reader, &t)
	
	for i := 0; i < t; i++ {
		arr := make([]int, 3)
		fmt.Fscan(reader, &arr[0], &arr[1], &arr[2])
		sort.Ints(arr)

		fmt.Fprintf(writer, "Scenario #%v:\n", i+1)
		
		if arr[0]*arr[0]+arr[1]*arr[1] == arr[2]*arr[2] {
			fmt.Fprintf(writer, "%v\n\n", "yes")
			continue
		}
		fmt.Fprintf(writer, "%v\n\n", "no")
	}
}
