package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   moneys := []int{1, 5, 10, 20, 50, 100}
	arr := make([]int, 6)
   max := 0
   idx := -1
   cnt := 0
	for i := 0; i < len(arr); i++ {
      fmt.Fscan(reader, &arr[i])
      temp := arr[i]*moneys[i]
      if max < temp || (max == temp && cnt > arr[i]) {
         max = temp
         idx = i
         cnt = arr[i]
      }
	}

   fmt.Fprintf(writer, "%v\n", moneys[idx])
}
