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
   
   var a, b string
   fmt.Fscan(reader, &a, &b)

   if a == "null" {
      fmt.Fprintln(writer, "NullPointerException")
      fmt.Fprintln(writer, "NullPointerException")
      return
   }
   if b == "null" {
      fmt.Fprintln(writer, "false")
      fmt.Fprintln(writer, "false")
      return
   }

   if a == b {
      fmt.Fprintln(writer, "true")
      fmt.Fprintln(writer, "true")
      return
   }
   fmt.Fprintln(writer, "false")
   if len(a) != len(b) {
      fmt.Fprintln(writer, "false")
      return
   }

   for i, _ := range a {
      if a[i] == b[i] {
         continue
      }
      if strings.ToUpper(string(a[i])) != string(b[i]) && strings.ToUpper(string(b[i])) != string(a[i]) {
         fmt.Fprintln(writer, "false")
         return
      }
   }
   fmt.Fprintln(writer, "true")
}
