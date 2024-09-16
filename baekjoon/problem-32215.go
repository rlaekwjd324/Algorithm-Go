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

   var n, m, k int
   fmt.Fscan(reader, &n, &m, &k)

   fmt.Fprintln(writer, m*k+m)
}
