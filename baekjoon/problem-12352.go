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

   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var n int
      fmt.Fscan(reader, &n)
      
      arr := make([]float64, n)
      for j := 0; j < n; j++ {
         var h float64
         fmt.Fscan(reader, &h)

         arr[j] = h
      }
      
      for j := 1; j < n-1; j++ {
         s := (arr[j-1]+arr[j+1])/2.0
         if s >= arr[j] {
            continue
         }
         arr[j] = s
      }

      fmt.Fprintf(writer, "Case #%v: %.6f\n", i+1, arr[n-2])
   }
}
