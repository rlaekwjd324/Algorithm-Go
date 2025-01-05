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

   pre := 0
   for i := 0; i < 4; i++ {
      sum := 0
      for j := 0; j < 4; j++ {
         var n int
         fmt.Fscan(reader, &n)

         sum += n
      }
      if i == 0 {
         pre = sum
      } else {
         if pre != sum {
            fmt.Fprintln(writer, "not magic")
            return
         }
      }
   }

   fmt.Fprintln(writer, "magic")
}
