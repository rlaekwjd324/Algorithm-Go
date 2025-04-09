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
   
   str := "UAPC"
   a := ""
   for _, v := range str {
      if !strings.Contains(s, string(v)){
         a += string(v)
      }
   }
   fmt.Fprintln(writer, a)
}
