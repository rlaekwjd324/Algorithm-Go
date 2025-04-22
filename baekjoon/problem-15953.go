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
   
   bfV := []int{500, 300, 200, 50, 30, 10}
   bfC := []int{1, 2, 3, 4, 5, 6}
   afV := []int{512, 256, 128, 64, 32}
   afC := []int{1, 2, 4, 8, 16}

   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var a, b int
      fmt.Fscan(reader, &a, &b)
      
      sum := 0
      if a != 0 {
         for j := 0; j < len(bfV); j++ {
            a -= bfC[j]
            if a <= 0 {
               sum += bfV[j]
               break
            }
         }
      }
      if b != 0 {
         for j := 0; j < len(afV); j++ {
            b -= afC[j]
            if b <= 0 {
               sum += afV[j]
               break
            }
         }
      }
      fmt.Fprintf(writer, "%v\n", sum*10000)
   }
}
