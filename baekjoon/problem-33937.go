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

   var f, s string
   fmt.Fscan(reader, &f, &s)

   mo := "aeiou"

   fs := ""
   hasMo := false
   for i, v := range f {
      if !strings.Contains(mo, string(v)) {
         if hasMo {
            break
         }
      } else {
         hasMo = true
      }
      fs += string(v)
      if i == len(f)-1 {
         fs = ""
      }
   }
   ss := ""
   hasMo = false
   for i, v := range s {
      if !strings.Contains(mo, string(v)) {
         if hasMo {
            break
         }
      } else {
         hasMo = true
      }
      ss += string(v)
      if i == len(s)-1 {
         ss = ""
      }
   }
   if fs == "" || ss == "" {
      fmt.Fprintln(writer, "no such exercise")
      return
   }
   fmt.Fprintf(writer, "%v%v\n", fs, ss)
}
