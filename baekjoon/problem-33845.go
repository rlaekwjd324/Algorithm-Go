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
   
   var s, t string
   fmt.Fscan(reader, &s, &t)
   
   a := ""
   for _, v := range t {
      if !strings.Contains(s, string(v)) {
         a += string(v)
      }
   }

   fmt.Fprintf(writer, "%v\n", a)
}
