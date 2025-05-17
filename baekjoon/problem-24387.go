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
   
   aArr := make([]int, 3)
   bArr := make([]int, 3)

   for i := 0; i < 6; i++ {
      var n int
      fmt.Fscan(reader, &n)
      if i < 3 {
         aArr[i] = n
      } else {
         bArr[i-3] = n
      }
   }
   sort.Ints(aArr)
   sort.Ints(bArr)
   
   sum := aArr[0] * bArr[0] + aArr[1] * bArr[1] + aArr[2] * bArr[2]

   fmt.Fprintf(writer, "%v\n", sum)
}
