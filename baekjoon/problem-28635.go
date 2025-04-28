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
   
   var m, a, b int
   fmt.Fscan(reader, &m, &a, &b)
   
   if a > b {
      fmt.Fprintf(writer, "%v\n", m-a+b)
   }else {
      fmt.Fprintf(writer, "%v\n", b-a)
   }
}
