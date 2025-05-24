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

   var s, c, o, n int
   fmt.Fscan(reader, &s, &c, &o, &n)
   
   a := 0
   for {
      if s >= 2 {
         s -= 2
      } else if n >= 2 {
         n -= 2
      } else {
         break
      }
      if c >= 4 {
         c -= 4
      } else if o >= 2 {
         o -= 2
      } else {
         break
      }
      if o >= 1 {
         o -= 1
      } else if c >= 2 {
         c -= 2
      } else {
         break
      }
      if n >= 1 {
         n -= 1
      } else if s >= 1 {
         s -= 1
      } else {
         break
      }
      a++
   }

   fmt.Fprintf(writer, "%v\n", a)
}
