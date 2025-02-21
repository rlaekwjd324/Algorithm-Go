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
   
   var n string
   fmt.Fscan(reader, &n)
   
   fmt.Fprintln(writer, n)
   for i := 0; i < 25; i++ {
      a := ""
      for _, v := range n {
         newV := int(v)-1
         if newV < 65 {
            newV = 90
         }
         a += string(newV)
      }
      n = a
      a += "\n"
      fmt.Fprint(writer, a)
   }

}
