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
   
   arr := []int{1}
   sum := []int{1}
   for i := 1; i <= n; i++ {
      var x int
      fmt.Fscan(reader, &x)
      if x == 0 {
         fmt.Fprintf(writer, "%v\n", 0)
         continue
      }

      if len(sum) >= x {
         fmt.Fprintf(writer, "%v\n", sum[x-1])
         continue
      }
      for j := len(sum); j < x; j++ {
         arr = append(arr, arr[j-1]+j+1)
         sum = append(sum, sum[j-1]+arr[j])
      }
      
      fmt.Fprintf(writer, "%v\n", sum[x-1])
   }
}
