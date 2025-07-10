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
      var a, c int
      var r, g, b int
      fmt.Fscan(reader, &a, &c, &r, &g, &b)
      
      s1 := getValue(a, c, r+1, g, b)
      s2 := getValue(a, c, r, g+1, b)
      s3 := getValue(a, c, r, g, b+1)

      if s1 < s2 {
         if s2 < s3 {
            fmt.Fprintf(writer, "%v\n", "BLUE")
            continue
         }
         fmt.Fprintf(writer, "%v\n", "GREEN")
         continue
      }
      if s1 < s3 {
         fmt.Fprintf(writer, "%v\n", "BLUE")
         continue
      }
      fmt.Fprintf(writer, "%v\n", "RED")
   }
}

func getValue(a, c, r, g, b int) int {
   sum := a * (r*r+g*g+b*b)
   min := 16
   if r > g {
      if g > b {
         min = b
      }else {
         min = g
      }
   } else {
      if r > b {
         min = b
      }else {
         min = r
      }
   }
   sum += c*min
   return sum
}
