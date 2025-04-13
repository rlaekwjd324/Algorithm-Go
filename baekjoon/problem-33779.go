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
   
   var t string
   fmt.Fscan(reader, &t)
   
   for i := 0; i < len(t)/2; i++ {
      if t[i] != t[len(t)-i-1] {
         fmt.Fprintln(writer, "boop")
         return
      }
   }
   fmt.Fprintln(writer, "beep")
}
