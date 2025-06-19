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

   var a, b, c int
   fmt.Fscan(reader, &a, &b, &c)
   
   if a == 2015 {
      if b <= 2 {
         fmt.Fprintf(writer, "%v\n", 1)
         return
      }
      if b <= 5 {
         fmt.Fprintf(writer, "%v\n", 2)
         return
      }
      if b <= 8 {
         fmt.Fprintf(writer, "%v\n", 3)
         return
      }
      if b <= 11 {
         fmt.Fprintf(writer, "%v\n", 4)
         return
      }
      fmt.Fprintf(writer, "%v\n", 5)
      return
   }
   sum := 4
   sum += (a-2016)*4
   if b <= 2 {
      sum += 1
   } else if b <= 5 {
      sum += 2
   } else if b <= 8 {
      sum += 3
   } else if b <= 11 {
      sum += 4
   } else {
      sum += 5
   }
   fmt.Fprintf(writer, "%v\n", sum)
}
