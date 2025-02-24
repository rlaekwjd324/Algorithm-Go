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

   for i := 1; i <= t; i++ {
      fmt.Fprintf(writer, "Material Management %v\nClassification ---- End!\n", i)
   }
}
