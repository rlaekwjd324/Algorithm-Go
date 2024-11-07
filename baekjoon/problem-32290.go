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

   var l, r, x int
   fmt.Fscanln(reader, &l, &r, &x)
   
   arr := make([]int, r-l+1)
   for i := l; i <= r; i++ {
      or := i | x
      arr[i-l] = or
   }
   sort.Slice(arr, func(i, j int) bool {
      return arr[i] < arr[j]
   })
   for i := 0; i < arr[r-l]; i++ {
      if i != arr[i] {
         fmt.Fprintln(writer, i)
         return
      }
   }
   fmt.Fprintln(writer, arr[r-l]+1)
}
