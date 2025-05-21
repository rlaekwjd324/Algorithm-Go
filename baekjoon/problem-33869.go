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

   var w, s string
   fmt.Fscan(reader, &w, &s)

   w1 := ""
   for i := 0; i < len(w); i++ {
      if !strings.Contains(w1, string(w[i])) {
         w1 += string(w[i])
      }
   }
   for i := 0; i < 26; i++ {
      if !strings.Contains(w1, string(rune(65+i))) {
         w1 += string(rune(65+i))
      }
   }
   for i := 0; i < len(s); i++ {
      fmt.Fprintf(writer, "%v", string(w1[int(s[i])-65]))
   }
}
