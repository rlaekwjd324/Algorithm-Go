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

   var t, v int
   fmt.Fscan(reader, &t, &v)
   
   fmt.Fprintln(writer, t*v)
}
