package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n int
   fmt.Fscanln(reader, &n)
   
   answer := ""

   for i := 0; i < 3; i++ {
      line := ""
      if i == 0 {
         line += strings.Repeat("*", n)
         line += strings.Repeat("x", n)
         line += strings.Repeat("*", n)
      }else if i == 1 {
         line += strings.Repeat(" ", n)
         line += strings.Repeat("x", n)
         line += strings.Repeat("x", n)
      }else if i == 2 {
         line += strings.Repeat("*", n)
         line += strings.Repeat(" ", n)
         line += strings.Repeat("*", n)
      }
      answer += strings.Repeat(line+"\n", n)
   }
   fmt.Fprintln(writer, answer)
}
