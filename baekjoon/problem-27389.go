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

   var a float64
   fmt.Fscan(reader, &a)
   
   fmt.Fprintf(writer, "%.2f\n", a/4)
}
