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

	arr := make([]int, 4)
	fmt.Fscan(reader, &arr[0], &arr[1], &arr[2], &arr[3])

	sort.Ints(arr)
	min := float64(arr[0])/float64(arr[2])+float64(arr[1])/float64(arr[3])
	tmp := float64(arr[1])/float64(arr[2])+float64(arr[0])/float64(arr[3])
	if min > tmp {
		fmt.Fprintf(writer, "%v %v %v %v\n", arr[1], arr[2], arr[0], arr[3])
		return
	}
	fmt.Fprintf(writer, "%v %v %v %v\n", arr[0], arr[2], arr[1], arr[3])
}
