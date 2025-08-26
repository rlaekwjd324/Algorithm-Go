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

   var n int
   fmt.Fscan(reader, &n)
   
   var arr = []int{1}
   i := 0
   h := 1
   for {
      h += 2
      i++
      arr = append(arr, h*h+arr[i-1])
      if arr[i] > n {
         break
      }
   }

   fmt.Fprintf(writer, "%v\n", i)
}
