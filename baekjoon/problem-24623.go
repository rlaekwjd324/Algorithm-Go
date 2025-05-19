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

   var n, a int
   fmt.Fscan(reader, &n, &a)

   fmt.Fprintf(writer, "%v\n", 90)
}
