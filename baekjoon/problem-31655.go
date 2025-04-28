package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
   "strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()
   
   var s string
   fmt.Fscan(reader, &s)
   
   f, _ := strconv.Atoi(strings.Split(s, "/")[0])
   m, _ := strconv.Atoi(strings.Split(s, "/")[1])
   
   if f > 12 {
      fmt.Fprintln(writer, "EU")
      return
   }
   if m > 12 {
      fmt.Fprintln(writer, "US")
      return
   }
   fmt.Fprintln(writer, "either")
}
