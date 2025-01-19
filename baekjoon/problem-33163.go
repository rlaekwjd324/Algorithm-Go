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

   var n int
   var s string
   fmt.Fscan(reader, &n, &s)
   
   a := ""
   for _, v := range s {
      switch v {
      case 'J':
         a += "O"
      case 'O':
         a += "I"
      case 'I':
         a += "J"
      }
   }

   fmt.Fprintln(writer, a)
}
