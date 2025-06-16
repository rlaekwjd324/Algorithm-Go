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

   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var n int
      fmt.Fscan(reader, &n)
      
      min := 100000.0
      answer := 0.0 
      for j := 0; j < n; j++ {
         var w, v float64
         fmt.Fscan(reader, &w, &v)
         temp := v/w
         if temp == min {
            if answer > v {
               answer = v
            }
            continue
         } 
         if temp < min {
            min = temp
            answer = v
         }
      }
      fmt.Fprintf(writer, "%v\n", int(answer))
   }
}
