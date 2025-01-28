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

   var a, b int
   fmt.Fscan(reader, &a, &b)
   sum := a * 1000 + b * 10000

   fmt.Fprintln(writer, sum)
}
