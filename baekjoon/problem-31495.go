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

   var s string
   fmt.Fscan(reader, &s)

   if len(s) <= 2 {
     fmt.Fprintln(writer, "CE")
     return
   }
   for i, v := range s {
     if i == 0 || i == len(s) - 1 {
         if v != '"' {
             fmt.Fprintln(writer, "CE")
             return
         }
     }
   }
   fmt.Fprintln(writer, s[1:len(s)-1])
}
