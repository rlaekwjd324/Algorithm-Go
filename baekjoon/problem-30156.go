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
      var w string
      fmt.Fscan(reader, &w)
      
      c := 0
      for _, v := range w {
         if v == 'a' {
            c++
         }
      }
      if c >= len(w)/2 {
         fmt.Fprintln(writer, len(w)-c)
      } else {
         fmt.Fprintln(writer, c)
      }
   }
}
