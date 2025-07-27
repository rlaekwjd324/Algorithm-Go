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

   var d, m, w, i, j, k int64
   fmt.Fscan(reader, &d, &m, &w, &i, &j, &k)

   s := d*m*(k-1) + d*(j-1) + i-1
   dd := s % w

   fmt.Fprintf(writer, "%v\n", string(rune('a' + dd)))
}
