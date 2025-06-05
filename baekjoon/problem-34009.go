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

   var n int
   fmt.Fscan(reader, &n)
   
   arr := make([]int, n)
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      arr[i] = a
   }
   sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
   aSum := 0
   bSum := 0
   for i := 0; i < n; i++ {
      if (n-i) % 2 == 0 {
         aSum += arr[i]
      } else {
         bSum += arr[i]
      }
      if bSum >= aSum {
         fmt.Fprintln(writer, "Bob")
         return
      }
   }
   fmt.Fprintln(writer, "Alice")
}
