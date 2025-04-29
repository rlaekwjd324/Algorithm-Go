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
   
   var s string
   fmt.Fscan(reader, &s)
   
   c := len(s)
   for {
      if strings.Contains("aeiouAEIOU", string(s[c-1])) {
         fmt.Fprintf(writer, "%s%s\n", s[:c], "ntry")
         return
      }
      c--
      if c == 0 {
         break
      }
   }
   fmt.Fprintf(writer, "%s%s\n", s[:c], "ntry")
}
