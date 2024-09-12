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

   var n, m int
   fmt.Fscan(reader, &n, &m)

   count := m-n+1
   if n %2 == 0 {
      count++
   }
   if count %2 == 0 {
      count = count/2
   }else {
      count = (count+1)/2
   }

   fmt.Fprintln(writer, count)
}
