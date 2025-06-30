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

   var a1, a2, a3, b1, b2, b3, l int
   fmt.Fscan(reader, &a1, &a2, &a3, &b1, &b2, &b3, &l)
   
   var aArr, bArr []int
   if a1 < a2 {
      if a2 < a3 {
         aArr = []int{a3, a2, a1}
         bArr = []int{b3, b2, b1}
      } else {
         aArr = []int{a2, a3, a1}
         bArr = []int{b2, b3, b1}
      }
   } else {
      if a1 < a3 {
         aArr = []int{a3, a1, a2}
         bArr = []int{b3, b1, b2}
      } else {
         aArr = []int{a1, a3, a2}
         bArr = []int{b1, b3, b2}
      }
   }
   sum := 0
   c := 0
   for {
      if sum >= l {
         fmt.Fprintf(writer, "%v\n", c)
         return
      }
      c++
      if bArr[0] > 0 {
         sum += aArr[0]
         bArr[0]--
         continue
      }
      if bArr[1] > 0 {
         sum += aArr[1]
         bArr[1]--
         continue
      }
      if bArr[2] > 0 {
         sum += aArr[2]
         bArr[2]--
         continue
      }
      fmt.Fprintf(writer, "%v\n", 0)
      return
   }
}
