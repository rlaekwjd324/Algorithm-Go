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

   var t1, t2, t3, t4, t5, t6, br int
   fmt.Fscan(reader, &t1, &t2, &t3, &t4, &t5, &br, &t6)

   h := 0
   m := 10
   m += t3
   h += t4
   m += t5
   m += (br+1)*t6
   
   h = t1 - h
   m = t2 - m
   for {
      if m >= 0 {
         break
      }
      m += 60
      h -= 1
   }
   
   if h < 10 {
      fmt.Fprintf(writer, "0%v ", h)
   } else {
      fmt.Fprintf(writer, "%v ", h)
   }
   
   if m < 10 {
      fmt.Fprintf(writer, "0%v\n", m)
      return
   }
   fmt.Fprintf(writer, "%v\n", m)
}
