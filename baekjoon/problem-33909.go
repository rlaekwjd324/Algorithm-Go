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

   var s, c, o, n int
   fmt.Fscan(reader, &s, &c, &o, &n)
   
   

   fmt.Fprintf(writer, "%v\n", "")
}
