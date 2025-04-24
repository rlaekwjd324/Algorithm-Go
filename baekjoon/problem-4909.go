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
   
   for {
      sum := 0
      max := 0
      min := 10
      for i := 0; i<6; i++ {
         var a int
         fmt.Fscan(reader, &a)
         if max < a {
            max = a
         }
         if min > a {
            min = a
         }
         sum += a
      }
      if sum == 0 {
         break
      }
      sum = sum - min - max
      fmt.Fprintf(writer, "%v\n", float64(sum)/4.0)
   }
}
