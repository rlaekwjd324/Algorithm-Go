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

   var n string
   fmt.Fscan(reader, &n)
   
   t := "ns"
   m := "aeiou"
   
   if strings.Contains(t, string(n[len(n)-1])) || strings.Contains(m, string(n[len(n)-1])) {
      c := 0
      for i := len(n)-1; i >= 0; i-- {
         if strings.Contains(m, string(n[i])) {
            c++
            if c == 2 {
               fmt.Fprintf(writer, "%v\n", i+1)
               return
            }
         }
      }
   } else {
      c := 0
      for i := len(n)-1; i >= 0; i-- {
         if strings.Contains(m, string(n[i])) {
            c++
            if c == 1 {
               fmt.Fprintf(writer, "%v\n", i+1)
               return
            }
         }
      }
   }
   fmt.Fprintf(writer, "%v\n", -1)
}
